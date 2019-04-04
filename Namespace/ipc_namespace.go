package main

import(
	"log"
	"os"
	"os/exec"
	"syscall"
)


//IPC Namespace 用来隔离System V IPC 和POSIX message quques。每一个IPC Namespace
//都有自己的System V IPC 和 POSIX message queue
//System V 引用了三种高级进程间的通信机制:消息队列，共享内存和信号量
//IPC对象(消息队列,共享内存和信号量)存在于内核中而不是文件系统中，由用户控制释放，不像管道的释放又
//由内核控制IPC对象通过其标识符引用和访问，所有IPC对象内核空间有唯一性的标志ID，在用户空间的唯一性标志位key
//LINUX IPC 继承自 System IPC

//REF :https://blog.csdn.net/qq_38211852/article/details/80475818
func main(){

	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscanll.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err !=  nil{
		log.Fatal(err)
	}

}