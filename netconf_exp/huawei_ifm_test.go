/**
 * @Author: DPY
 * @Description: 接口管理测试
 * @File:  huawei_ifm_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/9 18:14
 */

package netconf_exp

import (
	"encoding/xml"
	"testing"
)

type Ifm struct {
	XMLName        xml.Name    `xml:"ifm"`
	Xmlns          string      `xml:"xmlns,attr"`
	FormatVersion  string      `xml:"format-version,attr"`
	ContentVersion string      `xml:"content-version,attr"`
	Interfaces     []Interface `xml:"interfaces>interface"`
}

type Interface struct {
	IfName         string `xml:"ifName"`                     // 接口名称
	IfIndex        string `xml:"ifIndex"`                    // 接口索引,1-134217727
	IfClass        string `xml:"ifClass"`                    // 主接口还是子接口,mainInterface,subInterface
	IfPhyType      string `xml:"ifPhyType"`                  // 接口类型
	IfPosition     string `xml:"ifPosition"`                 // 接口位置,1/0/1
	IfParentIfName string `xml:"ifParentIfName"`             // 如果是子接口显示父接口名称
	IfDescr        string `xml:"ifDescr"`                    // 接口描述
	IfTrunkIfName  string `xml:"ifTrunkIfName"`              // 如果被eth-trunk绑定，显示trunk接口名称
	IsL2SwitchPort string `xml:"isL2SwitchPort"`             // 是否二层交换口,true,false
	IfAdminStatus  string `xml:"ifAdminStatus"`              // 接口管理状态,up,down
	IfOperStatus   string `xml:"ifDynamicInfo>ifOperStatus"` // 接口操作状态,up,down
	IfPhyStatus    string `xml:"ifDynamicInfo>ifPhyStatus"`  // 接口物理状态,up,down
	IfLinkStatus   string `xml:"ifDynamicInfo>ifLinkStatus"` // 接口链路状态,up,down
	//VrfName         string       `xml:"vrfName"`                       // 接口vrf名称
	IfServiceType string `xml:"ifServiceType"` // 接口属性,None,TrunkMember,StackMember,FabricMember
	IfMtu         string `xml:"ifMtu"`         // 接口mtu
	IfMac         string `xml:"ifMac"`         // 接口mac地址
	//L2SubIfFlag     string       `xml:"l2SubIfFlag"`                   // 是否二层子接口,true,false
	AddrType        string    `xml:"ifmAm4>addrCfgType"`            // 接口地址类型,config,unnumbered
	AddrUnNumIfName string    `xml:"ifmAm4>unNumIfName"`            // 如果接口是借用地址,显示借用的接口名称
	AddrContent     []CfgAddr `xml:"ifmAm4>am4CfgAddrs>am4CfgAddr"` // 接口IP地址列表
}

type CfgAddr struct {
	IfIpAddr   string `xml:"ifIpAddr"`   // ipv4地址,192.168.52.209
	SubnetMask string `xml:"subnetMask"` // ipv4地址掩码,255.255.255.252
	AddrType   string `xml:"addrType"`   // 地址类型,main,sub,unnumber
}

// TestIgmGet 获取接口信息
func TestIgmGet(t *testing.T) {
	var info Ifm
	info.Interfaces = []Interface{{IfName: "MEth0/0/0"}}
	info.Xmlns = "http://www.huawei.com/netconf/vrp"
	info.FormatVersion = "1.0"
	info.ContentVersion = "1.0"
	xmlBytes, err := makeGetXml(info)
	if err != nil {
		t.Fatal(err)
	}
	session, err := NewSession(Target)
	if err != nil {
		t.Fatal(err)
	}
	reply, err := session.Exec(NCRule(xmlBytes))
	if err != nil {
		t.Fatal(err)
	}
	var rData Ifm
	var r = Result{Data: &rData}

	err = xml.Unmarshal([]byte(reply.Data), &r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetIndentJson(rData))
}

// TestUnmarshalData 测试解析接口数据到结构体
func TestUnmarshalData(t *testing.T) {
	var data = `<data>
            <ifm xmlns="http://www.huawei.com/netconf/vrp" format-version="1.0" content-version="1.0">
              <interfaces>
                <interface>
                  <ifName>MEth0/0/0</ifName>
                  <ifIndex>4</ifIndex>
                  <ifPhyType>MEth</ifPhyType>
                  <ifPosition>0/0/0</ifPosition>
                  <ifParentIfName></ifParentIfName>
                  <ifDescr/>
                  <ifTrunkIfName/>
                  <isL2SwitchPort>false</isL2SwitchPort>
                  <ifAdminStatus>up</ifAdminStatus>
                  <ifMtu>1500</ifMtu>
                  <ifMac>c8b6-d3aa-56ea</ifMac>
                  <ifServiceType>None</ifServiceType>
                  <ifClass>mainInterface</ifClass>
                  <ifDynamicInfo>
                    <ifOperStatus>down</ifOperStatus>
                    <ifPhyStatus>down</ifPhyStatus>
                    <ifLinkStatus>down</ifLinkStatus>
                  </ifDynamicInfo>
                  <ifmAm4>
                    <unNumIfName></unNumIfName>
                    <addrCfgType>config</addrCfgType>
                    <am4CfgAddrs>
                      <am4CfgAddr>
                        <ifIpAddr>192.168.0.1</ifIpAddr>
                        <subnetMask>255.255.255.0</subnetMask>
                        <addrType>main</addrType>
                      </am4CfgAddr>
                    </am4CfgAddrs>
                  </ifmAm4>
                </interface>
              </interfaces>
            </ifm>
          </data>`
	var r Result
	var rz Ifm
	r.Data = &rz
	err := xml.Unmarshal([]byte(data), &r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(GetIndentJson(rz))
}
