// @Author: Perry
// @Date  : 2020/1/20
// @Desc  :

package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
	"testing"
)

type Config struct {
	AppName  string `ini:"app_name"`
	LogLevel string `ini:"log_level"`

	MySQL MySQLConfig `ini:"mysql"`
	Redis RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

/*将配置文件映射到结构体*/
func TestIniV3_1(t *testing.T) {
	cfg, err := ini.Load("ini_3.ini")
	if err != nil {
		fmt.Println("load my.ini failed: ", err)
	}

	c := Config{}
	cfg.MapTo(&c)

	fmt.Printf("%+v\n",c)
}
