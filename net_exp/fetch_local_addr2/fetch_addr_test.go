// @Author: Perry
// @Date  : 2020/1/17
// @Desc  : 通过dial udp来获取本地发起访问的本地IP地址

package fetch_local_addr2

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func Test1(t *testing.T) {
	localIp := GetOutboundIP()
	fmt.Println(localIp)
}
