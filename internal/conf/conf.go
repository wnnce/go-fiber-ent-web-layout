package conf

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Bootstrap struct {
	Server *Server `json:"server" yaml:"server"`
	Data   *Data   `json:"data" yaml:"data"`
	Jwt    *Jwt    `json:"jwt" yaml:"jwt"`
}

type Server struct {
	Host    string        `json:"host" yaml:"host"`
	Port    uint          `json:"port" yaml:"port"`
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}

type Data struct {
	Database struct {
		Driver   string `json:"driver" yaml:"driver"`
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		DbName   string `json:"dbName" yaml:"db-name"`
	} `json:"database" yaml:"database"`
	Redis struct {
		Host        string        `json:"host" yaml:"host"`
		Port        int           `json:"port" yaml:"port"`
		Index       int           `json:"index" yaml:"index"`
		Username    string        `json:"username" yaml:"username"`
		Password    string        `json:"password" yaml:"password"`
		ReadTimeout time.Duration `json:"readTimeout" yaml:"read-timeout"`
		WireTimeout time.Duration `json:"wireTimeout" yaml:"wire-timeout"`
	} `json:"redis" yaml:"redis"`
}

type Jwt struct {
	Issue      string        `json:"issue" yaml:"issue"`
	ExpireTime time.Duration `json:"expireTime" yaml:"expire-time"`
	Secret     string        `json:"secret" yaml:"secret"`
}

// ReadConfig 读取配置文件
// path 配置文件路径
func ReadConfig(path string) *Bootstrap {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(dir + path)
	if err != nil {
		panic(err)
	}
	bootstrap := &Bootstrap{}
	err = yaml.Unmarshal(file, bootstrap)
	if err != nil {
		panic(err)
	}
	return bootstrap
}
