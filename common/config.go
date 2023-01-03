package common

import (
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

const CONFIG_FILE = "config.yaml"

var Conf Config

func GetConfig() *Config {
	return &Conf
}

type (
	DNSConfig struct {
		Domain  string `yaml:"domain"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Timeout int    `yaml:"timeout"`
		TTL     int    `yaml:"ttl"`
	}
	MySQLConfig struct {
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Name    string `yaml:"name"`
		Charset string `yaml:"charset"`
	}
	RedisConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Pass string `yaml:"pass"`
		Db   int    `yaml:"db"`
	}
	ServerConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	JWTConfig struct {
		Secret  string `yaml:"secret"`
		Expires int    `yaml:"expires"`
	}
	WeChatConfig struct {
		AppID     string `yaml:"appid"`
		AppSecret string `yaml:"secret"`
		Token     string `yaml:"token"`
		AesKey    string `yaml:"aeskey"`
		WeChatDb  int    `yaml:"wxdb"`
	}
	QiniuConfig struct {
		Access string `yaml:"access"`
		Secret string `yaml:"secret"`
		Bucket string `yaml:"bucket"`
	}
	QQConfig struct {
		Appid  string `yaml:"appid"`
		Secret string `yaml:"secret"`
		Token  string `yaml:"token"`
	}
	MailConfig struct {
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
	}
)

type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Server ServerConfig `yaml:"server"`
	DNS    DNSConfig    `yaml:"dns"`
	Jwt    JWTConfig    `yaml:"jwt"`
	WeChat WeChatConfig `yaml:"wechat"`
	Qiniu  QiniuConfig  `yaml:"qiniu"`
	QQ     QQConfig     `yaml:"qq"`
	Mail   MailConfig   `yaml:"mail"`
}

var (
	// ConfigName -
	ConfigName = "config.yaml"
)

func defaultConfig() *Config {
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

	conf.DNS.Host = getEnvDefault("DNS_HOST", "127.0.0.1")
	conf.DNS.Domain = getEnvDefault("DNS_DOMAIN", "example.com")
	conf.DNS.Port, _ = strconv.Atoi(getEnvDefault("DNS_PORT", "53"))
	conf.DNS.Timeout, _ = strconv.Atoi(getEnvDefault("DNS_TIMEOUT", "5"))
	conf.DNS.TTL, _ = strconv.Atoi(getEnvDefault("DNS_TTL", "600"))

	return conf
}

// TemplateConfig show config template
func TemplateConfig() (data []byte, err error) {
	// yamlData
	data, err = yaml.Marshal(defaultConfig())
	return
}

// LoadConfig -
func LoadConfig() (err error) {
	var yf []byte
	_, err = os.Stat(CONFIG_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			Conf = *defaultConfig()
			return nil
		}
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
