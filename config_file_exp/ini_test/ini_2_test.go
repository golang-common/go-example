// @Author: Perry
// @Date  : 2020/1/20
// @Desc  : 

package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"testing"
)

/*简单的ini文件读取测试*/
func TestIniV2_1(t *testing.T) {
	cfg, err := ini.Load("ini_2.ini")
	if err != nil {
		log.Fatal("fail to read file: ", err)
	}

	fmt.Println("app name:", cfg.Section("").Key("app_name").String())
	fmt.Println("log level:", cfg.Section("").Key("log_level").String())

	fmt.Println("mysql ip:", cfg.Section("mysql").Key("ip").String())
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
	fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
	fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
	fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())

	fmt.Println("Redis IP:", cfg.Section("redis").Key("ip").String())
	redisPort, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis Port:", redisPort)
}

/*写入配置文件*/
func TestIniV2_2(t *testing.T) {
	cfg := ini.Empty()

	defaultSection := cfg.Section("")
	defaultSection.NewKey("app_name", "awesome web")
	defaultSection.NewKey("log_level", "DEBUG")

	mysqlSection, err := cfg.NewSection("mysql")
	if err != nil {
		fmt.Println("new mysql section failed:", err)
		return
	}
	mysqlSection.NewKey("ip", "127.0.0.1")
	mysqlSection.NewKey("port", "3306")
	mysqlSection.NewKey("user", "root")
	mysqlSection.NewKey("password", "123456")
	mysqlSection.NewKey("database", "awesome")

	redisSection, err := cfg.NewSection("redis")
	if err != nil {
		fmt.Println("new redis section failed:", err)
		return
	}
	redisSection.NewKey("ip", "127.0.0.1")
	redisSection.NewKey("port", "6381")

	err = cfg.SaveTo("ini_2_save.ini")
	if err != nil {
		fmt.Println("SaveTo failed: ", err)
	}

	err = cfg.SaveToIndent("ini_2_save_pretty.ini", "\t")
	if err != nil {
		fmt.Println("SaveToIndent failed: ", err)
	}

	cfg.WriteTo(os.Stdout)
	fmt.Println()
	cfg.WriteToIndent(os.Stdout, "\t")
}
