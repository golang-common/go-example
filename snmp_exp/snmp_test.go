/**
 * @Author: daipengyuan
 * @Description:
 * @File:  snmp_test
 * @Version: 1.0.0
 * @Date: 2021/3/13 16:43
 */

package snmp_exp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"testing"
)

func TestSNMP1(t *testing.T) {
	snmp := gosnmp.Default
	snmp.Target = "172.21.0.52"
	snmp.Version = gosnmp.Version2c
	snmp.Community ="daipengyuan@123"
	fmt.Println("stt")
	snmp.Timeout = 200000000
	err := snmp.Connect()
	if err != nil {
		t.Fatal(err)
	}
	pdus, err := snmp.BulkWalkAll(".1.3.6.1.2.1.1.2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pdus)
}
