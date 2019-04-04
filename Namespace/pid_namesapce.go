package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

//PID Namespace是用来隔离进程ID的。同样一个进程在不同的PID Namespace里
//是可以拥有不同的PID的。这样就可以理解，在docker container里面，使用ps -ef经常会发现，
// 在容器内,前台运行的那个进程PID是1.

func main() {

	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
