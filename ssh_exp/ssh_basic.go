// @Author: Perry
// @Date  : 2020/1/15
// @Desc  : 

package ssh_exp

import (
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

const (
	sshNetTimeout = 300 * time.Millisecond
)

/*新建SSH Client*/
func NewSSHClient(ip, port, user, passwd string) (*ssh.Client, error) {
	socket := ip + ":" + port
	// 建立网络连接
	conn, err := net.DialTimeout("tcp", socket, sshNetTimeout)
	// ssh客户端协商
	sshConfig := getClientConfig(user, passwd)                         // 获取ssh
	cc, chans, reqs, err := ssh.NewClientConn(conn, socket, sshConfig) // 发起ssh client协商
	if err != nil {
		if cc != nil {
			_ = cc.Close()
		}
		return nil, err
	}
	return ssh.NewClient(cc, chans, reqs), nil
}

/*获取ssh配置*/
func getClientConfig(user, password string) *ssh.ClientConfig {
	var (
		sshConfig ssh.Config
	)
	sshConfig.SetDefaults()
	sshConfig.Ciphers = append(sshConfig.Ciphers, "aes128-ctr", "aes192-ctr", "aes256-ctr",
		"aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc")
	clientConfig := &ssh.ClientConfig{
		Config:          sshConfig,
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         sshNetTimeout,
	}
	return clientConfig
}
