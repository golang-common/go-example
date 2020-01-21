// @Author: Perry
// @Date  : 2020/1/14
// @Desc  : 

package main

import (
	"bytes"
	"dpy/exp/ssh_exp"
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	msgSeparator             = " ## ## "
	shellReadTimeout         = 5 * time.Second
	shellNewSessionSleepTime = 300 * time.Millisecond
)

/*新建SSH Session*/
func NewShellSession(client *ssh.Client) (*ShellPipeSession, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("dumb", 200, 200, modes)
	if err != nil {
		return nil, err
	}
	sessionPipe, err := setShellPipe(session)
	if err != nil {
		_ = session.Close()
		return nil, err
	}
	time.Sleep(shellNewSessionSleepTime)

	return sessionPipe, nil
}

/*设置ssh会话的输入输出管道*/
func setShellPipe(session *ssh.Session) (*ShellPipeSession, error) {
	w, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}
	r, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = session.Shell()
	if err != nil {
		return nil, err
	}
	return &ShellPipeSession{session, r, w}, nil
}

type ShellPipeSession struct {
	session *ssh.Session
	io.Reader
	io.WriteCloser
}

func (d *ShellPipeSession) Close() (err error) {
	err = d.session.Close()
	d.session = nil
	return
}

func (d *ShellPipeSession) ExecCmd(cmdList []string) (out []byte, err error) {
	var (
		outBytes = new([]byte)
		e        error
		wg       sync.WaitGroup
		cmdBytes = []byte("\n" + strings.Join(cmdList, "\n") + "\n" + msgSeparator + "\n")
	)
	_, err = d.Write(cmdBytes)
	if err != nil {
		return
	}
	wg.Add(1)
	go func() {
		e = d.readDelimit([]byte(msgSeparator), outBytes)
		wg.Done()
	}()
	err = d.waitTimeout(&wg)
	out = *outBytes
	if err != nil {
		return
	}
	return out, e
}

func (d *ShellPipeSession) readDelimit(sep []byte, rstBytes *[]byte) error {
	var out bytes.Buffer
	buf := make([]byte, 4096)
	pos := 0
	for {
		n, err := d.Read(buf[pos : pos+(len(buf)/2)])
		if err != nil {
			if n == 0 {
				out.Write(buf[0:pos])
			}
			*rstBytes = out.Bytes()
			return err
		}
		if n > 0 {
			outBytes := append(out.Bytes(), buf[0:pos+n]...)
			if index := bytes.Index(outBytes, sep); index > -1 {
				*rstBytes = outBytes[0:index]
				return nil
			}
			if pos > 0 {
				out.Write(buf[0:pos])
				copy(buf, buf[pos:pos+n])
			}
			pos = n
		}
	}
}

func (d *ShellPipeSession) waitTimeout(wg *sync.WaitGroup) error {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	select {
	case <-ch:
		return nil
	case <-time.After(shellReadTimeout):
		return errors.New("read ssh result timeout ")
	}
}

func Test1(t *testing.T) {
	client, err := ssh_exp.NewSSHClient("127.0.0.1", "22", "daipengyuan", "Dpy,./883675")
	if err != nil {
		log.Fatal(err)
	}
	pipeSession, err := NewShellSession(client)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := pipeSession.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	content, err := pipeSession.ExecCmd([]string{"ls"})
	if err != nil {
		log.Println(err)
	}
	log.Println(string(content))
	content, err = pipeSession.ExecCmd([]string{"pwd"})
	if err != nil {
		log.Println(err)
	}
	log.Println(string(content))
}
