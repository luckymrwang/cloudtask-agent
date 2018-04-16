package cache

import "github.com/cloudtask/libtools/gounits/system"

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createCommandFile(directory string, cmd string) (string, error) {

	fname := "run.cmd"
	fpath := directory + "/" + fname
	if ret := system.FileExist(fpath); ret {
		if err := os.Remove(fpath); err != nil {
			return "", err
		}
	}

	body := fmt.Sprintf("@echo off\r\ncd /d %s\r\n%s", `%~dp0`, cmd)
	if err := ioutil.WriteFile(fpath, []byte(body), 0777); err != nil {
		return "", err
	}
	return fname, nil
}
