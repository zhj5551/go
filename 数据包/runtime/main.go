package main

import (
	"fmt"
	"runtime"
)

func init() {
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：", runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}

func main() {
	/*
		1. NumCPU：返回当前系统的 CPU 核数量
		2. GOMAXPROCS：设置最大的可同时使用的 CPU 核数
		3. Gosched：让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行,这个函数的作用是让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，Go 会自动地把与该 goroutine 处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞。
		4. Goexit：退出当前 goroutine(但是defer语句会照常执行)
		5. NumGoroutine：返回正在执行和排队的任务总数。runtime.NumGoroutine函数在被调用后，会返回系统中的处于特定状态的Goroutine的数量。这里的特指是指Grunnable\Gruning\Gsyscall\Gwaition。处于这些状态的Groutine即被看做是活跃的或者说正在被调度。
			注意：垃圾回收所在Groutine的状态也处于这个范围内的话，也会被纳入该计数器。
		6. GOOS：目标操作系统
		7. runtime.GC：会让运行时系统进行一次强制性的垃圾收集
			1.强制的垃圾回收：不管怎样，都要进行的垃圾回收。
			2.非强制的垃圾回收：只会在一定条件下进行的垃圾回收（即运行时，系统自上次垃圾回收之后新申请的堆内存的单元（也成为单元增量）达到指定的数值）。
		8. GOROOT：获取goroot目录
		9. GOOS : 查看目标操作系统 很多时候，我们会根据平台的不同实现不同的操作，就而已用GOOS了：
	*/
	//获取goroot目录：
	fmt.Println("GOROOT-->", runtime.GOROOT())

	//获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS) // GOOS--> darwin，mac系统
}
