package tempexp

import (
	"net"
	"testing"
)

func TestIp(t *testing.T) {
	_, netw, err := net.ParseCIDR("ff::ff01/126")
	if err != nil {
		t.Fatal(err)
	}
	ones, _ := netw.Mask.Size()
	bp := ones / 8
	bm := ones % 8
	if bm == 0 {
		bp -= 1
		bm = 8
	}
	ip := netw.IP
	t.Log(ip.String())
	t.Log(bp)
	t.Log(bm)
	t.Log(1 << (8 - bm))
	ip[bp] = ip[bp] + 1<<(8-bm)
	t.Log(ip.String())
}

func ipAddrPlus1(ip net.IP) net.IP {
	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i] == 255 {
			ip[i] = 0
			continue
		}
		ip[i] += 1
		break
	}
	return ip
}

func TestIpCon(t *testing.T) {
	ip := ipAddrPlus1(net.ParseIP("254.255.255.255").To4())
	t.Log(ip)
}

type ss struct {
	a int
	b int
}

func TestMol(t *testing.T) {
	var sl = []ss{{a: 1, b: 11}, {a: 2, b: 22}}
	for _, v := range sl {
		s := v
		s.a = 11
		t.Logf("%+v", s)
	}
	t.Logf("%+v",sl)
}

func TestSli(t *testing.T){
	var a = []int{0}
	t.Log(a[:len(a)-1])
}