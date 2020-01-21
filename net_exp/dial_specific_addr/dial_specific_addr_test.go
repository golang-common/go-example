// @Author: Perry
// @Date  : 2020/1/19
// @Desc  : 测试网络连接

package dial_specific_addr

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

/*测试指定IPv4源地址连接*/
func TestIPv4Src(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "192.168.160.203:0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := net.Dialer{LocalAddr: addr, Timeout: 3 * time.Second}
	conn, err := d.Dial(addr.Network(), "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(conn.LocalAddr().String())
	_ = conn.Close()
}

/*测试简单IPv6连接*/
func TestIPv6(t *testing.T) {
	conn, err := net.Dial("udp6", "[::1]:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(conn.LocalAddr().Network())
	fmt.Println(conn.LocalAddr().String())
	_ = conn.Close()
}

/*测试指定IPv6源地址连接*/
func TestIPv6Src(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "[::1]:0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dialer := net.Dialer{LocalAddr: addr, Timeout: 3 * time.Second}
	conn, err := dialer.Dial(addr.Network(), "[::1]:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(conn.LocalAddr().Network())
	fmt.Println(conn.LocalAddr().String())
}
