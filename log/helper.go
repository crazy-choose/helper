package log

import (
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

func homeDir() string {
	u, e := user.Current()
	if e != nil {
		return "./"
	}
	return u.HomeDir
}

func exeName() string {
	var en string
	//获取执行文件名字
	exeFile, err := exec.LookPath(os.Args[0])
	if err != nil {
		return time.Now().String()
	} else {
		begin := strings.LastIndex(exeFile, "/") + 1
		end := len(exeFile)
		en = exeFile[begin:end]
	}
	return en
}

func defaultName() string {
	return homeDir() + "/logs/" + exeName()
}
