package gostrip

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/goretk/gore"
	"github.com/pkg/errors"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

const regex = "a-f0-9"

type Opts struct {
	IsForce bool
}

func Strip(filename string, opts *Opts) error {
	fn, err := filepath.Abs(filename)
	if err != nil {
		return errors.WithStack(err)
	}
	_f, err := gore.Open(fn)
	if err != nil {
		return errors.WithStack(err)
	}
	defer _f.Close()
	logger.SuccessF("开始混淆: %s", filename)
	// 读取文件
	buf, err := os.ReadFile(fn)
	if err != nil {
		return errors.WithStack(err)
	}
	f := &GoFile{GoFile: _f, Origin: buf}
	err = f.Init()
	if err != nil {
		return errors.WithStack(err)
	}

	// 获取编译器信息
	cmp, err := f.GetCompilerVersion()
	if err != nil {
		return errors.WithStack(err)
	}
	logger.SuccessF("编译器: %s %s %s", cmp.Name, cmp.SHA, cmp.Timestamp)

	// 混淆文件名
	for _, t := range TypeStringOffsets.FileName {
		f.SetBytes(t.Offset, t.Length, utils.RandomBytesByLength(t.Length, regex))
	}
	logger.Success("混淆文件名")

	// 混淆函数名
	for _, t := range TypeStringOffsets.Func {
		f.SetBytes(t.Offset, t.Length, utils.RandomBytesByLength(t.Length, regex))
	}
	logger.Success("混淆函数名")

	mod := f.BuildInfo.ModInfo

	// 以下为暴力替换，可能会存在不知名 Bug
	if opts.IsForce {
		// TODO: 使用更优雅的方法
		typs, err := f.GetTypes()
		if err == nil {
			for _, t := range typs {
				if t.PackagePath == "" || gore.IsStandardLibrary(t.PackagePath) || len(t.Name) > 50 {
					continue
				}
				if t.Kind == reflect.Struct &&
					strings.HasPrefix(t.PackagePath, mod.Main.Path) {
					logger.Debug(t.Name)
					f.Replace([]byte(t.Name), utils.RandomBytesByLength(len(t.Name), regex), -1)
				}
			}
			logger.Warn("混淆结构体名称")
		}
	}

	// 混淆ModInfo
	repls := [][]byte{[]byte(mod.Path), []byte(mod.Main.Path), []byte(mod.Main.Version)}
	for _, d := range mod.Deps {
		repls = append(repls, []byte(d.Path), []byte(d.Sum)[3:])
	}
	for _, repl := range repls {
		f.Replace(repl, utils.RandomBytesByLength(len(repl), regex), -1)
	}
	logger.Success("混淆ModInfo")

	// 混淆BuildID
	f.Replace([]byte(f.BuildID), utils.RandomBytesByLength(len(f.BuildID), regex), -1)
	logger.Success("混淆BuildID")

	// // 混淆版本 - 直接替换
	// f.Replace([]byte(cmp.Name), utils.RandomBytesByLength(len(cmp.Name), regex), -1)
	// logger.Success("混淆版本")

	newName := fmt.Sprintf("%s_massup", filename)
	if err := f.Save(newName); err != nil {
		return errors.WithStack(err)
	}
	logger.SuccessF("混淆成功，新文件: %s", newName)

	return nil
}
