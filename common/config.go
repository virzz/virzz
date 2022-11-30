package common

import (
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

const CONFIG_FILE = "config.yaml"

var Conf Config

// Config Hongyan Config
type Config struct {
	MySQL struct {
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Name    string `yaml:"name"`
		Charset string `yaml:"charset"`
	} `yaml:"mysql"`
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Pass string `yaml:"pass"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	// Move
	Jwt struct {
		Secret  string `yaml:"secret"`
		Expires int    `yaml:"expires"`
	} `yaml:"jwt"`
	WeChat struct {
		AppID     string `yaml:"appid"`
		AppSecret string `yaml:"secret"`
		Token     string `yaml:"token"`
		AesKey    string `yaml:"aeskey"`
		WeChatDb  int    `yaml:"wxdb"`
	} `yaml:"wechat"`
	Qiniu struct {
		Access string `yaml:"access"`
		Secret string `yaml:"secret"`
		Bucket string `yaml:"bucket"`
	} `yaml:"qiniu"`
	QQ struct {
		Appid  string `yaml:"appid"`
		Secret string `yaml:"secret"`
		Token  string `yaml:"token"`
	} `yaml:"qq"`
	Mail struct {
		User string `yaml:"user"`
		Stmp struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
			Pass string `yaml:"pass"`
		} `yaml:"stmp"`
		AliDirect struct {
			Region string `yaml:"region"`
			Access string `yaml:"access"`
			Secret string `yaml:"secret"`
		} `yaml:"alidirect"`
	} `yaml:"mail"`
}

var (
	// ConfigName -
	ConfigName = "config.yaml"
)

// TemplateConfig show config template
func TemplateConfig() []byte {
	conf := &Config{}

	conf.Jwt.Expires, _ = strconv.Atoi(getEnvDefault("JWT_EXPIRES", "3600"))
	conf.Jwt.Secret = getEnvDefault("JWT_SECRET", "jwt_secret_bfbf094bf0049f6e4b79d098dc2b363c")

	conf.MySQL.Charset = "utf8mb4"
	conf.MySQL.Host = getEnvDefault("MYSQL_HOST", "mariadb")
	conf.MySQL.User = getEnvDefault("MYSQL_USER", "virzz")
	conf.MySQL.Pass = getEnvDefault("MYSQL_PASS", "virzz9999")
	conf.MySQL.Name = getEnvDefault("MYSQL_NAME", "virzz_platform")

	conf.Redis.Host = getEnvDefault("REDIS_HOST", "127.0.0.1")
	conf.Redis.Port, _ = strconv.Atoi(getEnvDefault("REDIS_PORT", "6379"))
	conf.Redis.Db, _ = strconv.Atoi(getEnvDefault("REDIS_DB", "0"))

	conf.Server.Host = getEnvDefault("SERVER_HOST", "127.0.0.1")
	conf.Server.Port, _ = strconv.Atoi(getEnvDefault("SERVER_PORT", "9999"))

	// yamlData
	data, err := yaml.Marshal(conf)
	if err != nil {
		return []byte(err.Error())
	}
	return data
}

// LoadConfig -
func LoadConfig() (err error) {
	var yf []byte
	_, err = os.Stat(CONFIG_FILE)
	if err != nil && os.IsNotExist(err) {
		return err
	}
	if yf, err = os.ReadFile(CONFIG_FILE); err != nil {
		return err
	}
	if err = yaml.Unmarshal(yf, &Conf); err != nil {
		return err
	}
	return nil
}
