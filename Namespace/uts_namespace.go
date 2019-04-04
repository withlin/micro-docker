package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

//UTS Namespace 主要用来隔离nodename和domainname的两个系统标志。在UTS Namespace里面
//每个Namespace允许有自己的hostname
func main() {
	//用来指定被fork出来的新进程内的初始命令，默认用sh来执行
	cmd := exec.Command("sh")
	//调用系统参数
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
