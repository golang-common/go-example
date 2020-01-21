// @Author: Perry
// @Date  : 2020/1/15
// @Desc  : 

package sftp_exp

import (
	"bufio"
	"bytes"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"os"
)

func NewSftpSession(client *ssh.Client) (*SftpSession, error) {
	session, err := sftp.NewClient(client)
	if err != nil {
		return nil, err
	}
	return &SftpSession{session}, nil
}

type SftpSession struct {
	session *sftp.Client
}

func (d *SftpSession) Close() (err error) {
	err = d.session.Close()
	d.session = nil
	return
}

/*发送文件内容content到目录path*/
func (d *SftpSession) Send(path string, content []byte) (size int, err error) {
	fd, err := d.session.Create(path)
	if err != nil {
		return
	}
	w := bufio.NewWriter(fd)
	size, err = w.Write(content)
	return
}

/*从对端filePath接收文件*/
func (d *SftpSession) Receive(filePath string) (content []byte, size int64, err error) {
	var out = new(bytes.Buffer)
	fd, err := d.session.Open(filePath)
	if err != nil {
		return
	}
	r := bufio.NewReader(fd)
	size, err = r.WriteTo(out)
	if err != nil {
		return
	}
	content = out.Bytes()
	return
}

/*获取当前目录*/
func (d *SftpSession) GetWD() (wd string, err error) {
	return d.session.Getwd()
}

/*获取dirPath目录下的文件列表*/
func (d *SftpSession) ReadDir(dirPath string) (fList []os.FileInfo, err error) {
	return d.session.ReadDir(dirPath)
}

/*获取filePath指向的文件信息,判断文件是否存在*/
func (d *SftpSession) Stat(filePath string) (stat os.FileInfo, err error) {
	return d.session.Stat(filePath)
}
