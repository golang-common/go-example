// @Author: Perry
// @Date  : 2020/1/13
// @Desc  : 读取配置文件及默认值

package ini_test

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"testing"
)

func TestIniV1(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(dir)
	cfg, err := ini.Load("/Users/daipengyuan/code/go/src/dpy/exp/config_file_exp/ini_test/ini_1.ini")
	if err != nil {
		fmt.Printf("fail to read file: %v", err)
		os.Exit(1)
	}

	// 正常读取全局值
	fmt.Println("App Mode:", cfg.Section(ini.DefaultSection).Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// 读取有限制的值
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	fmt.Println("Email Protocol:",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))
}
