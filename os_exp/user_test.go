/**
 * @Author: DPY
 * @Description:
 * @File:  user_test
 * @Version: 1.0.0
 * @Date: 2021/11/10 16:03
 */

package os_exp

import (
	"os/user"
	"testing"
)

//user.Current()
//user.Lookup()
//user.LookupId()
//user.LookupGroup()
//user.LookupGroupId()

func TestCurrent(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", u)
}

func TestLookup(t *testing.T) {
	u, err := user.Lookup("lyonsdpy")
	if err != nil {
		t.Fatal(err)
	}
	u.GroupIds()
	t.Logf("%+v\n", u)
}

func TestLookupId(t *testing.T) {
	u, err := user.LookupId("501")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", u)
}

func TestLookupGroup(t *testing.T) {
	g, err := user.LookupGroup("staff")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", g)
}

func TestLookupGroupId(t *testing.T) {
	g, err := user.LookupGroupId("20")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", g)
}

