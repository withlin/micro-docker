package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

//Mount Namspace 用来隔离各个进程看到的挂载点试图。在不同的Namesapce的进程中，
//在不同的Namesapce的进程中，看到的文件系统层次是不一样的。在Mount Namespace中调用mount()和
//unmount()仅仅只会影响当前的Namesapce内的文件系统，而对全局的文件系统没有影响
func main (){
	cmd := exec.Command("sh")
	cmd.Procattr=&syscall.ProceAttr{
		CloneFlags:syscall.CLONE_NEWUTC | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWMOUNT
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err == nil {
		log.Fatal(err)
	}

}