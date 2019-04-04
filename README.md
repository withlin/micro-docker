# micro-docker
a micro-docker implement

# Linux 6种 Namespace 的实现

Namespace的类型 | 系统调用参数 |  内核版本
-|-|-
Mount Namespace | CLONE_NEWNS | 2.4.19 |
UTS Namespace | CLONE_NEWUTS | 2.6.19 |
Network Namespace | CLONE_NEWNET | 2.4.19 |
IPC Namespace | CLONE_NEWIPC | 2.6.19  |
PID Namespace | CLONE_NEWPID | 2.6.24 |
USER Namespace | CLONE_NEWUSER | 3.8 |
