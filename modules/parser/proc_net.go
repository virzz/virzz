package parser

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/downloader"
)

var TcpState = map[string]string{
	"01": "ESTABLISHED",
	"02": "SYN-SENT",
	"03": "SYN-RECEIVED",
	"04": "FIN-WAIT-1",
	"05": "FIN-WAIT-2",
	"06": "TIME-WAIT",
	"07": "CLOSED",
	"08": "CLOSE-WAIT",
	"09": "LAST-ACK",
	"0A": "LISTEN",
	"0B": "CLOSING",
	"0C": "NEW-SYN-RECEIVED",
}

type ProcNetTcp struct {
	LocalIP    string
	LocalPort  string
	RemoteIP   string
	RemotePort string
	State      string
}

func hexToIP(hexStr string) string {
	ip, err := hex.DecodeString(hexStr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v.%v.%v.%v", ip[3], ip[2], ip[1], ip[0])
}

func hexToPort(hex string) string {
	port, err := basic.HexToDec(hex)
	if err != nil {
		return ""
	}
	return port
}

// parseProcNetTcp Parse /proc/net/tcp
func parseProcNetTcp(src string) (string, error) {
	var data []byte
	var err error
	var dataStr string
	filename := path.Join(os.TempDir(), "tcp")
	if strings.HasPrefix(src, "http") {
		if err = downloader.SigleFetch(src, filename); err != nil {
			return "", err
		}
		data, err = os.ReadFile(filename)
		if err != nil {
			return "", err
		}
		dataStr = string(data)
	} else if fs, err := os.Stat(src); err == nil && !fs.IsDir() {
		data, err = os.ReadFile(src)
		if err != nil {
			return "", err
		}
		dataStr = string(data)
	} else {
		dataStr = src
	}
	// parse
	result := make([]ProcNetTcp, 0)
	re := regexp.MustCompile(`(?m)(\d:) ([0-9a-fA-F]+):([0-9a-fA-F]+) ([0-9a-fA-F]+):([0-9a-fA-F]+) ([0-9a-fA-F]+)`)
	for _, match := range re.FindAllStringSubmatch(dataStr, -1) {
		if len(match) > 6 {
			state := ""
			if v, ok := TcpState[match[6]]; ok {
				state = v
			}
			result = append(result, ProcNetTcp{
				LocalIP:    hexToIP(match[2]),
				LocalPort:  hexToPort(match[3]),
				RemoteIP:   hexToIP(match[4]),
				RemotePort: hexToPort(match[5]),
				State:      state,
			})
		}
	}
	tableCells := make([][]*simpletable.Cell, 0)
	for _, v := range result {
		tableCells = append(tableCells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%s:%s", v.LocalIP, v.LocalPort)},
			{Text: fmt.Sprintf("%s:%s", v.RemoteIP, v.RemotePort)},
			{Align: simpletable.AlignCenter, Text: v.State},
		})
	}
	return common.TableOutputV2(
		tableCells,
		// Header
		[]*simpletable.Cell{
			{Text: "Local"},
			{Text: "Remote"},
			{Align: simpletable.AlignCenter, Text: "State"},
		},
		// Footer
		[]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 3, Text: fmt.Sprintf("Connections - Total: %d", len(tableCells))},
		},
	), nil
}
