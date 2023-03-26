package tongji

import (
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/httpreq"
)

var once utils.OncePlus

// UniqID 生成设备ID = sha1(AppName + Hostname MAC)[:16]
func UniqID(name string) string {
	hostname, _ := os.Hostname()
	mac := ""
	i, err := net.InterfaceByIndex(0)
	if err == nil {
		mac = string(i.HardwareAddr)
	}
	md5Str, _ := hash.Sha1Hash([]byte(fmt.Sprintf("%s-%s-%s", name, mac, hostname)))
	return md5Str[:16]
}

// Tongji 使用统计
// 程序名/版本号/操作系统/架构/语言/设备ID
func Tongji(url, name, ver string) {
	if os.Getenv("VIRZZ_NO_TONGJI") != "" {
		return
	}
	tj := viper.New()
	tj.SetConfigFile(path.Join("$HOME", ".config", "virzz", "tongji.yaml"))
	if tj.GetBool("tongji") {
		return
	}
	logger.Info("Init tongji at the first time")
	once.Do(func() error {
		_, err := httpreq.R().
			SetQueryParams(map[string]string{
				"name": name,
				"ver":  ver,
				"os":   runtime.GOOS,
				"arch": runtime.GOARCH,
				"ln":   strings.Split(os.Getenv("LANG"), ".")[0],
				"id":   UniqID(name),
				// "t":    strconv.Itoa(int(time.Now().Unix())),
			}).Get(url)
		return err
	})
	tj.Set("tongji", true)
	tj.WriteConfig()
	tj.SafeWriteConfig()
}
