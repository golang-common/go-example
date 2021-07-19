/**
 * @Author: daipengyuan
 * @Description:
 * @File:  basic_test
 * @Version: 1.0.0
 * @Date: 2021/6/15 10:07
 */

package nftables_exp

import (
	"fmt"
	"net"
	"testing"
)

func TestRune(t *testing.T) {
	a := []rune{'w', 'K', 'g', 'B', 'C', 'g', '=', '='}
	fmt.Println(a)
}

func TestAddr(t *testing.T) {
	addr, _ := NewIPAddr("192.168.1.1/24")
	t.Log(getInverseMask(getMask(*addr.Mask, len(addr.IP))))
	t.Log([]byte(computeGapRange(addr)))
}

type IPAddr struct {
	*net.IPAddr
	CIDR bool
	Mask *uint8
}

func getIP(ip *IPAddr) []byte {
	if !ip.IsIPv6() {
		return ip.IP.To4()
	}
	return ip.IP.To16()
}

func (ip *IPAddr) IsIPv6() bool {
	if ip.IP.To4() == nil {
		return true
	}
	return false
}

func getMask(ml uint8, l int) []byte {
	mask := make([]byte, l)
	fullBytes := ml / 8
	leftBits := ml % 8
	for i := 0; i < int(fullBytes); i++ {
		mask[i] = 0xff
	}
	if leftBits != 0 {
		m := uint8(0x80)
		v := uint8(0x00)
		for i := 0; i < int(leftBits); i++ {
			v += m
			m = m >> 1
		}
		mask[fullBytes] ^= v
	}

	return mask
}

func computeGapRange(e1 *IPAddr) net.IP {
	imask1 := getInverseMask(getMask(*e1.Mask, len(e1.IP)))
	bip1 := addInverseMaskPlusOne(getIP(e1), imask1)

	return net.IP(bip1)
}

func getInverseMask(mask []byte) []byte {
	inv := make([]byte, len(mask))
	for i := 0; i < len(mask); i++ {
		inv[i] = ^mask[i]
	}

	return inv
}

func addInverseMaskPlusOne(ip, mask []byte) []byte {
	r := make([]byte, len(ip))
	for i := 0; i < len(mask); i++ {
		r[i] = ip[i] | mask[i]
	}
	for i := len(r) - 1; i >= 0; i-- {
		r[i]++
		if r[i] != 0 {
			return r
		}
	}

	return r
}

func NewIPAddr(addr string) (*IPAddr, error) {
	if _, ipnet, err := net.ParseCIDR(addr); err == nil {
		// Found a valid CIDR address
		ones, _ := ipnet.Mask.Size()
		mask := uint8(ones)
		return &IPAddr{
			&net.IPAddr{
				IP: ipnet.IP,
			},
			true,
			&mask,
		}, nil
	}
	// Check if addr is just ip address in a non CIDR format
	ip := net.ParseIP(addr)
	if ip == nil {
		return nil, fmt.Errorf("%s is invalid ip address", addr)
	}
	mask := uint8(32)
	if ip.To4() == nil {
		mask = uint8(128)
	}
	_, ipnet, err := net.ParseCIDR(addr + "/" + fmt.Sprintf("%d", mask))
	if err != nil {
		return nil, err
	}
	return &IPAddr{
		&net.IPAddr{
			IP: ipnet.IP,
		},
		true,
		&mask,
	}, nil
}
