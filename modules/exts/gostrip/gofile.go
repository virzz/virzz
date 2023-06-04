package gostrip

import (
	"bytes"
	"debug/gosym"
	"os"
	"reflect"
	"sync"
	"unsafe"

	"github.com/goretk/gore"
	"github.com/pkg/errors"
)

type TypeData struct {
	Offset uint64
	Length uint64
}

type TypeStringOffset struct {
	Func     []TypeData
	FileName []TypeData
	Types    []TypeData
}

var TypeStringOffsets *TypeStringOffset = new(TypeStringOffset)

type GoFile struct {
	*gore.GoFile
	Origin       []byte
	InitPackages sync.Once
}

func (f *GoFile) SetBytes(offset uint64, length uint64, value []byte) error {
	valOff := int(length) - len(value)
	if valOff > 0 {
		value = append(value, make([]byte, valOff)...)
	}
	for i := uint64(0); i < length; i++ {
		f.Origin[offset+i] = value[i]
	}
	return nil
}

func (f *GoFile) GetBytes(offset uint64, length uint64) []byte {
	return f.Origin[offset : offset+length]
}

func (f *GoFile) Replace(origin []byte, new []byte, n int) {
	f.Origin = bytes.Replace(f.Origin, origin, new, n)
}

func (f *GoFile) Save(filename string) error {
	err := os.WriteFile(filename, f.Origin, 0755)
	return err
}

func (f *GoFile) getPCLNTabAddr() uint64 {
	vv := reflect.ValueOf(f.GoFile).Elem().
		FieldByName("fh").Elem().Elem().
		FieldByName("file").Elem()
	// PE
	if vv.FieldByName("pclntabAddr").IsValid() {
		return vv.FieldByName("pclntabAddr").Uint()
	}
	sections := vv.FieldByName("Sections")
	j := 0
	for j = 0; j < sections.Len(); j++ {
		section := sections.Index(j).Elem()
		name := section.FieldByName("Name").String()
		if name == "__gopclntab" ||
			name == ".gopclntab" || name == ".data.rel.ro.gopclntab" {
			return section.FieldByName("Addr").Uint()
		}
	}
	return 0
}

//lint:ignore U1000 todo
func (f *GoFile) getModuledataSectionAddr() uint64 {
	vv := reflect.ValueOf(f.GoFile).Elem().
		FieldByName("fh").Elem().Elem().
		FieldByName("file").Elem().
		FieldByName("Sections")
	for j := 0; j < vv.Len(); j++ {
		section := vv.Index(j).Elem()
		name := section.FieldByName("Name").String()
		addr := section.FieldByName("Addr").Uint()
		if name == ".noptrdata" || name == ".data" || name == "__noptrdata" {
			return addr
		}
	}
	return 0
}

func (f *GoFile) getFva(off uint64) (uint64, error) {
	vv := reflect.ValueOf(f.GoFile).Elem().
		FieldByName("fh").Elem().Elem().
		FieldByName("file").Elem().
		FieldByName("Sections")
	j := 0
	for j = 0; j < vv.Len(); j++ {
		section := vv.Index(j).Elem()
		addr := section.FieldByName("Addr").Uint()
		size := section.FieldByName("Size").Uint()
		offset := section.FieldByName("Offset").Uint()
		if addr <= off && off < addr+size {
			return off - addr + offset, nil
		}
	}
	return 0, gore.ErrSectionDoesNotExist
}

func (f *GoFile) Init() error {
	var returnVal error
	f.InitPackages.Do(func() {
		tab, err := f.PCLNTab()
		if err != nil {
			returnVal = err
			return
		}

		// // moduledata
		// md, err := f.Moduledata()
		// if err != nil {
		// 	returnVal = err
		// 	return
		// }
		// typeLink, err := md.TypeLink()
		// if err != nil {
		// 	returnVal = err
		// 	return
		// }
		// addr := md.Types().Address
		// for _, off := range typeLink {
		// 	logger.Debug(uint64(off) + addr)
		// }

		// xxx, _ := f.getFva(md.Types().Address)
		// logger.Debug(xxx)
		// logger.Debug(md.Types().Address)
		// d, _ := md.Types().Data()
		// logger.Debug(string(d))
		// logger.Debug(moduledata)
		// moduledataAddr := f.getModuledataSectionAddr()
		// logger.DebugF("moduledataAddr: %d", moduledataAddr)

		//混淆源码路径
		v := reflect.ValueOf(*tab)
		go12 := *(*gosym.LineTable)(unsafe.Pointer(v.FieldByName("go12line").Pointer()))
		v = reflect.ValueOf(&go12).Elem()
		DataLength := len(v.FieldByName("Data").Bytes())
		offsetFilePath, _ := f.getFva(f.getPCLNTabAddr())

		// 处理文件的结构
		fileMap := v.FieldByName("fileMap").MapRange()
		filetab := len(v.FieldByName("filetab").Bytes())
		offset := offsetFilePath + uint64(DataLength-filetab)
		for fileMap.Next() {
			TypeStringOffsets.FileName = append(
				TypeStringOffsets.FileName, TypeData{
					Offset: offset + fileMap.Value().Uint(),
					Length: uint64(len(fileMap.Key().String())),
				})
		}
		//混淆函数名称
		funcnametab := len(v.FieldByName("funcnametab").Bytes())
		funcNamesIter := v.FieldByName("funcNames").MapRange()
		offset = offsetFilePath + uint64(DataLength-funcnametab)
		for funcNamesIter.Next() {
			TypeStringOffsets.Func = append(TypeStringOffsets.Func, TypeData{
				Offset: offset + funcNamesIter.Key().Uint(),
				Length: uint64(len(funcNamesIter.Value().String())),
			})
		}
	})

	if returnVal != nil {
		return errors.WithStack(returnVal)
	}
	return nil
}
