package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMenoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d", syscall.Getegid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m -1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	cmd := exec.Command("proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflag: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		//得到fork出来映射在外部命名空间的pid
		fmt.Println("%v", cmd.Process.Pid)

		//在系统默认创建挂在memory subsystem的Hierarchy上创建cgroup
		os.Mkdir(path.Join(cgroupMenoryHierarchyMount, "testmemorylimit"), 0775)
		//将容器进程加入这个cgroup中
		ioutil.WriteFile(path.Join(cgroupMenoryHierarchyMount, "testmemorylimit",
			"tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		//限制cgroup进程使用
		ioutil.WriteFile(path.Join(cgroupMenoryHierarchyMount, "testmemorylimit",
			"memory.limit_in_bytes"), []byte("100m"), 0644)
		cmd.Process.Wait()

	}
}
