package log

import (
	"os"
	"os/exec"
	"strings"
)

func chToInstallDir() {
	//获取执行文件名字
	exeFile, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	} else {
		begin := strings.LastIndex(exeFile, "/") + 1
		end := len(exeFile)
		execName = exeFile[begin:end]
	}
}
