package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

//Network Namspace 用来隔离网络设备，IP地址端口等网络栈的Namespace。NetWork
//Namespace可以让每个容器拥有自己独立的(虚拟的)网路设备，而且容器内的应用可以绑定
//到自己的端口，每个Namesapce内的端口都不会互相冲突。在宿主机上搭建网桥后，就能
//很方便地实现容器之间的通信，而且不同的容器上的应用可以使用相同的端口


func main (){
	cmd := exec.Command("sh")
	cmd.ProcAttr=&syscall.ProceAttr{
		CloneFlags:syscall.CLONE_NEWUTC | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err == nil {
		log.Fatal(err)
	}

}