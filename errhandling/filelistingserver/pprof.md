关于pprof性能检测工具web命令无法查看图形化效果

需要安装Graphviz 

Graphviz 的windows安装教程
_https://blog.csdn.net/weixin_42654444/article/details/82108055_

总结：

1.安装Graphviz（_https://graphviz.org/download/_）

2.配置环境变量（Graphviz安装目录下的`bin`，以及`bin\dot.exe`）

3.restart

如果按以上教程安装之后任然无法使用web命令，需要执行一下步骤：

1. 打开cmd运行‘`dot -c`’命令安装依赖插件

2. 然后restart

性能测试常用命令：
`go test -bench . -cpuprofile cpu.out`

http网页测试，30秒内的CPU使用率：
`go tool pprof http://localhost:8888/debug/pprof/profile`
得到结果之后，执行
`web`
查看图形化效果

30秒内的内存使用率：
`go tool pprof http://localhost:8888/debug/pprof/heap`
`
更多命令查看pprof源码注释

example：web.go