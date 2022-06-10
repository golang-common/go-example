/**
 * @Author: DPY
 * @Description:
 * @File:  netconf.go
 * @Version: 1.0.0
 * @Date: 2022/6/9 17:17
 */

package netconf_exp

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
)

type target struct {
	host     string
	username string
	password string
}

var Target = target{
	host:     "172.22.100.59:830",
	username: "xsyxadmin",
	password: "Admin@1234",
}

type GetXml struct {
	XMLName xml.Name     `xml:"get"`
	Filter  GetFilterXml `xml:"filter"`
}

type GetFilterXml struct {
	XMLName xml.Name    `xml:"filter"`
	Type    string      `xml:"type,attr"`
	Data    interface{} `xml:",innerxml"`
}

type Result struct {
	XMLName xml.Name    `xml:"data"`
	Data    interface{} `xml:",any"`
}

type NCRule string

func (r NCRule) MarshalMethod() string {
	return string(r)
}

func makeGetXml(data interface{}) ([]byte, error) {
	var req GetXml
	req.Filter.Type = "subtree"
	req.Filter.Data = data
	r, err := xml.Marshal(req)
	return r, err
}

// GetIndentJson 优雅的使用JSON格式打印结构体数据
func GetIndentJson(obj interface{}) string {
	ret, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(ret)
}

// NewSession 新建netconf连接
func NewSession(target target) (*netconf.Session, error) {
	var sshConfig ssh.Config
	sshConfig.SetDefaults()
	sshConfig.Ciphers = append(sshConfig.Ciphers, "aes128-ctr", "aes192-ctr", "aes256-ctr",
		"aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc")
	sshConfig.KeyExchanges = append(
		sshConfig.KeyExchanges,
		"diffie-hellman-group-exchange-sha256",
		"diffie-hellman-group-exchange-sha1",
	)
	clientCfg := &ssh.ClientConfig{
		Config:          sshConfig,
		User:            target.username,
		Auth:            []ssh.AuthMethod{ssh.Password(target.password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	netconf.DefaultCapabilities = append(netconf.DefaultCapabilities,
		"http://www.huawei.com/netconf/capability/action/1.0",
		"http://www.huawei.com/netconf/capability/execute-cli/1.0",
		"http://www.huawei.com/netconf/capability/discard-commit/1.0")
	return netconf.DialSSH(target.host, clientCfg)
}
