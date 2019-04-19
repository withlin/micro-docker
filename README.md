# micro-docker
a micro-docker implement

## Linux 6种 Namespace 的实现

Namespace的类型 | 系统调用参数 |  内核版本
-|-|-
Mount Namespace   | CLONE_NEWNS   | 2.4.19  |
UTS Namespace     | CLONE_NEWUTS  | 2.6.19  |
Network Namespace | CLONE_NEWNET  | 2.4.19  |
IPC Namespace     | CLONE_NEWIPC  | 2.6.19  |
PID Namespace     | CLONE_NEWPID  | 2.6.24  |
USER Namespace    | CLONE_NEWUSER | 3.8     |


## Cgroup

* cgroup是对进程分组管理的一种机制，一个cgroup包含一组进程，并且可以在这个cgroup
  上增加Linux subsystem的各种配置，讲一组进程和一组subsystem的系统参数关联起来。

* subsystem 是一组资源控制的模块，一般包含如下几项：
  * blkio 设置对块设备(比如硬盘) 输入输出的访问控制
  * cpu 设置cgroup中进程的CPU被调度策略。
  * cpuacct 可以统计从cgroup中进程CPU占用。
  * cpuset 在多核机器上设设置cgroup中进程可以使用CPU和内存(此处内存仅仅使用于NUMA啊架构)。
  * devices 控制cgroup中进程对设备的访问。
  * freezer 用于挂起(suspend)和恢复(resume) cgroup中的进程。
  * memory 用于控制cgroup中进程的内存占用。
  * net_cls 用于讲cgroup中进程产生的网络包分类，以便Linux的tc (traffic controller) 可以
    根据分类区分出来某个cgroup的包并做限流或者监控。
  * net_prio 设置cgroup中进程产生的网络流量的优先级
  * ns 这个subsystem比较特殊，它的作用是使cgroup中的进程在新的Namesapce中fork新进程(NEWNS)时，
    创建出一个新的cgroup，这个cgroup包含新的Namesapce中的进程

* hierarchy的功能是把一组cgroup串成一个树状结构，一个这样的树状结构便是一个hierarchy
  通过这种树状结构，Cgroup可以做到继承。比如，系统对一组定时的任务进程通过cgroup1限制了
  CPU的使用率，然后其中有一个定时dump日志的进程还需要限制磁盘IO，为避免限制磁盘IO之后影
  响其他进程，就可以创建cgroup2，使其继承了cgroup1并且限制磁盘IO，这样cgroup2便继承了
  cgroup1中对CPU使用率的限制，并且增加了磁盘IO的限制而不影响到cgroup1中的其他进程。


* 系统在创建新的hierarchy之后，系统所有的进程都会加入这个hierarchy的cgroup根节点，
  这个cgroup根节点是hierarchy默认创建的。

* 一个subsystem只能附加到一个hierarchy上面。

* 一个heirarchy可以附加多个subsystem。

* 一个进程可以当作为多个cgroup的成员，但是这些cgroup必须在不同的hierarchy中。

* 一个进程fork出子进程时，子进程和父进程在同一个cgroup中，也可以根据需要将其移动
  到其他cgroup中。