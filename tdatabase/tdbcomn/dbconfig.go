package tdbcomn

import (
	"errors"
	"gopkg.in/ini.v1"
	"strings"
)

const DB_CONFIG_FILE = "tingo_db.ini"

type DBConfig struct {
	Driver   string `ini:"driver"`   //驱动名称
	Host     string `ini:"host"`     //ip或域名
	Port     int    `ini:"port"`     //端口
	User     string `ini:"username"` //用户名
	Password string `ini:"password"` //密码

	MaxIdleConns int `ini:"max_idle_conns"` //最大空闲连接数
	MaxLifeTime  int `ini:"max_life_ms"`    //一个连接最多存活时间,单位:ms
	MaxOpenConns int `ini:"max_conns"`      //最大连接数
}

func LoadDBConfig(driver string) (map[string]*DBConfig, error) {
	//忽略大小写
	//config, err := ini.InsensitiveLoad(DB_CONFIG_FILE)
	config, err := ini.Load(DB_CONFIG_FILE)

	if err != nil {
		return nil, err
	}
	result := make(map[string]*DBConfig, 2)
	sections := config.Sections()
	for _, sec := range sections {
		if strings.EqualFold(sec.Name(), "default") {
			continue
		}
		conf := &DBConfig{Host: "127.0.0.1", Port: 3306, User: "root"}
		err := sec.MapTo(conf)
		if err != nil {
			return nil, err
		}
		if conf.Driver == "" {
			panic(errors.New("not set driver: " + sec.Name()))
		}
		if conf.Host == "" {
			panic(errors.New("not set host: " + sec.Name()))
		}
		if conf.Port == 0 {
			panic(errors.New("not set port: " + sec.Name()))
		}

		if !strings.EqualFold(conf.Driver, driver) {
			continue
		}
		result[sec.Name()] = conf
	}
	return result, nil
}
