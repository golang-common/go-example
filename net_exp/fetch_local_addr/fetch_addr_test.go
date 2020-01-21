// @Author: Perry
// @Date  : 2020/1/2
// @Desc  : 获取所有接口的IP地址

package fetch_local_addr

import (
	"log"
	"net"
	"testing"
)

func Ips() (map[string]string, error) {
	ips := make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			ips[byName.Name] += v.String() +"  "
		}
	}
	return ips, nil
}

func Test1(t *testing.T) {
	mp, err := Ips()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range mp {
		log.Printf("ifname=%-5s\tifAddr=%10s\n", k, v)
	}
}
