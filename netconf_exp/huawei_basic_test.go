/**
 * @Author: DPY
 * @Description: netconf示例1
 * @File:  netconf1_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/9 17:12
 */

package netconf_exp

import (
	"strings"
	"testing"
)

// TestConnection 获取netconf会话，并打印会话id
func TestConnection(t *testing.T) {
	session, err := NewSession(Target)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	t.Log(session.SessionID)
}

// TestGetCap 获取并打印netconf协商的能力集
func TestGetCap(t *testing.T) {
	session, err := NewSession(Target)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	t.Log(strings.Join(session.ServerCapabilities, "\n"))
}
