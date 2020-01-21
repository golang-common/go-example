// @Author: Perry
// @Date  : 2020/1/17
// @Desc  : 

package basic_exp

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"testing"
)

func TestMemory(t *testing.T) {
	m, err := mem.VirtualMemory()
	if err != nil {
		t.Log(err)
	}
	t.Logf("total = %d bytes", m.Total)
	t.Logf("free = %d bytes", m.Free)
	t.Logf("usage = %d bytes", int(m.UsedPercent))
}

func TestCPU(t *testing.T) {
	c, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	// cpu的基本信息
	for _, cpuInfo := range c {
		fmt.Println(cpuInfo.String())
	}
	// cpu利用率信息
	percent, _ := cpu.Percent(0, true)
	fmt.Printf("cpu percent = %v\n", percent)
}

func TestHost(t *testing.T) {
	inf, err := host.Info()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", *inf)
}


