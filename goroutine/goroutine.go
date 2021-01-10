package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
go 1.14之前 goroutine是非抢占式调度，只能goroutine自己主动释放控制权，其他goroutine才可执行
gp 1.14之后引入了基于系统信号的异步抢占调度 goroutine已经支持抢占式调度，基本可以看做线程来使用
*/

func main() {
	fun2()
}

func fun2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++ // 1.14之前程序会死在这，1.14之后的Go协程已经支持抢占式调度了
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}

func fun1() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// go run -race (goroutine.go) 检测数据访问冲突
		go func(i int) { // race condition!
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
