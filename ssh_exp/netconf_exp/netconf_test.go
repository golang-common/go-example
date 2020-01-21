// @Author: Perry
// @Date  : 2020/1/15
// @Desc  : 

package netconf_exp

import (
	"github.com/lyonsdpy/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
)

const (
	closeXml = "<close-session/>"
)

func NewNetconfSession(client *ssh.Client, caps ...string) (*NcSession, error) {
	session, err := netconf.NewSSHSession2(client, caps...)
	if err != nil {
		return nil, err
	}
	return &NcSession{session}, nil
}

type NcSession struct {
	session *netconf.Session
}

func (d *NcSession) Close() (err error) {
	_, err = d.session.Exec(NCRule(closeXml))
	err = d.session.Close()
	d.session = nil
	return
}

func (d *NcSession) Exec(xmlStr string) (reply *netconf.RPCReply, err error) {
	reply, err = d.session.Exec(NCRule(xmlStr))
	return
}

type NCRule string

func (r NCRule) MarshalMethod() string {
	return string(r)
}
