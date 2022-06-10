/**
 * @Author: DPY
 * @Description:
 * @File:  io_test
 * @Version: 1.0.0
 * @Date: 2021/11/10 00:17
 */

package io_exp

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestDirFS(t *testing.T) {
	fsys := os.DirFS("/Users/lyonsdpy/Coding/go/src/golang_exp")
	entry, err := fs.ReadDir(fsys, "io_exp")
	if err != nil {
		t.Fatal(err)
	}
	for _, et := range entry {
		t.Log(et.Name())
	}
}

func TestStat(t *testing.T) {
	fsys := os.DirFS("/Users/lyonsdpy/Coding/go/src/golang_exp")
	finfo, err := fs.Stat(fsys, "go.mod")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(finfo.Name())
	t.Log(finfo.IsDir())
	t.Log(finfo.ModTime())
	t.Log(finfo.Size())
}

func TestReadFile(t *testing.T) {
	fsys := os.DirFS(`/Users/lyonsdpy/Coding/go/src/golang_exp`)
	fbytes, err := fs.ReadFile(fsys, "io_exp/dpy.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fbytes))
}

func TestGlob(t *testing.T) {
	fsys := os.DirFS(`/Users/lyonsdpy/Coding/go/src/golang_exp/io_exp`)
	matchList, err := fs.Glob(fsys, "dpy*")
	if err != nil {
		t.Fatal(err)
	}
	for _, match := range matchList {
		t.Log(match)
	}
}

func TestSub(t *testing.T) {
	fsys := os.DirFS(`/Users/lyonsdpy/Coding/go/src/golang_exp`)
	fsysSub, err := fs.Sub(fsys, "io_exp")
	if err != nil {
		t.Fatal(err)
	}
	fbytes, err := fs.ReadFile(fsysSub, "dpy.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(fbytes))
}

func TestValidPath(t *testing.T) {
	t.Log(fs.ValidPath("Users/lyonsdpy/Coding/go/src"))
	t.Log(fs.ValidPath("/Users/lyonsdpy/Coding/go/src"))
	t.Log(fs.ValidPath("Users/lyonsdpy/Coding/go/src/"))
}

func TestWalkDir(t *testing.T) {
	fsys := os.DirFS(`/Users/lyonsdpy/Coding/go/src/golang_exp/io_exp`)
	fn := func(path string, d fs.DirEntry, err error) error {
		fmt.Println("---")
		if err != nil {
			return err
		}
		fmt.Println(path)
		fmt.Println(d.Name())
		fmt.Println(d.IsDir())
		return nil
	}
	err := fs.WalkDir(fsys, ".", fn)
	if err != nil {
		t.Fatal(err)
	}
}