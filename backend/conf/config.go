package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
)

var conf Config

type Config struct {
	Server     Server     `yaml:"server"`
	Logger     Logger     `yaml:"logger"`
	Redis      Redis      `yaml:"redis"`
	Mysql      Mysql      `yaml:"mysql"`
	URLShorten URLShorten `yaml:"url_shorten"`
}

type Server struct {
	Port string `yaml:"port"` // 服务端口
	Name string `yaml:"name"`
}

type Logger struct {
	Path  string `yaml:"log_path"`
	Level string `yaml:"log_level"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Mysql struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type URLShorten struct {
	Prefix string `yaml:"prefix"`
	Expire int64  `yaml:"expire"`
}

// 配置文件路径
const ConfigPath = "./conf/config_%s.yaml"

func InitConfig() {
	env := os.Getenv("MODE_ENV")
	if env == "" {
		env = "dev"
	}

	configPath := fmt.Sprintf(ConfigPath, env)
	hlog.Info("read config from ", configPath)
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(dataBytes, &conf); err != nil {
		panic("conf.yaml配置文件读取失败:" + err.Error())
	}

	readConfigFromEnv()
	hlog.Infof("config = %+v", conf)
}

func readConfigFromEnv() {
	conf.Redis.Password = os.Getenv(conf.Redis.Password)
	conf.Mysql.Password = os.Getenv(conf.Mysql.Password)
}

func GetConfig() Config {
	return conf
}

func GetLogger() Logger {
	return conf.Logger
}

func GetRedis() Redis {
	return conf.Redis
}

func GetMysql() Mysql {
	return conf.Mysql
}

func GetURLShorten() URLShorten {
    return conf.URLShorten
}

func TestInit() {
	// 绝对路径
	dirPath, _ := getProjectPath()
	filePath := "./conf/config_dev.yaml"
	configPath := filepath.Join(dirPath, filePath)
	hlog.Info("read config from ", configPath)
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(dataBytes, &conf); err != nil {
		panic("conf.yaml配置文件读取失败:" + err.Error())
	}
	readConfigFromEnv()
}

func getProjectPath() (string, error) {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 从当前工作目录向上遍历，寻找main.go文件
	for {
		info, err := os.Stat(filepath.Join(cwd, "main.go"))
		if err == nil && !info.IsDir() {
			// 找到main.go，返回当前目录作为项目路径
			return cwd, nil
		} else if !os.IsNotExist(err) {
			// 其他错误
			return "", err
		}

		// 如果没找到，尝试进入上一级目录
		cwd = filepath.Dir(cwd)
		if cwd == "/" || cwd == "" {
			// 如果到达根目录仍然没找到，返回错误
			return "", fmt.Errorf("无法找到项目根目录")
		}
	}
}
