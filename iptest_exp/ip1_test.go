/**
 * @Author: daipengyuan
 * @Description:
 * @File:  ip1_test
 * @Version: 1.0.0
 * @Date: 2021/8/5 14:12
 */

package iptest_exp

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

// TestPadding 为IP地址补0
func TestPadding(t *testing.T) {
	t.Log(paddingIp(net.ParseIP("192.168.255.1")))
	t.Log(paddingIp(net.ParseIP("fe80::c560:d09c:6644:1dfd")))
	t.Log(paddingIp(net.ParseIP("fe80::44ef:eaff:feba:715")))
	t.Log(net.ParseIP("fe80:0000:0000:0000:44ef:eaff:feba:0715"))
}

// paddingIp 将IP地址相关位补0,形成一个IP地址固定长度的字符串格式
// ipv4: 192.168.255.1 --> 192.168.255.001
// ipv6: fe80::44ef:eaff:feba:715 --> fe80:0000:0000:0000:44ef:eaff:feba:0715
func paddingIp(ip net.IP) string {
	var rstList []string
	if ip == nil {
		return ""
	}
	if ip.To4() == nil {
		v6List := strings.Split(ip.String(), ":")
		for _, v := range v6List {
			if v == "" {
				for i := 0; i <= 8-len(v6List); i++ {
					rstList = append(rstList, "0000")
				}
				continue
			}
			rstList = append(rstList, fmt.Sprintf("%04s", v))
		}
		return strings.Join(rstList, ":")
	}
	v4List := strings.Split(ip.To4().String(), ".")
	for _, v := range v4List {
		rstList = append(rstList, fmt.Sprintf("%03s", v))
	}
	return strings.Join(rstList, ".")
}

