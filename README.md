# 第一章：基础知识
## 1.1 开发环境搭建
1. 下载安装go语言
下载地址： https://golang.google.cn/dl/
![image](https://user-images.githubusercontent.com/46108355/125374780-4bcbc600-e3ba-11eb-8317-4326a5431313.png)
- bin目录：包含了可执行程序，注意是可执行的，不需要解释执行。
- pkg目录：包含了使用的包或者说库。
- src目录：里面包含了go的代码源文件，其中仍按包的不同进行组织。
2. 配置vscode环境
随便点开一个 go 文件，在你的右下角会提示要你安装一些工具，安装的包有些由于墙的原因，无法下载，为了保证下载顺利，可以设置一下代理。
```
go env -w GOPROXY=https://goproxy.cn,direct
```
4. 配置环境变量
## 1.2 数据类型：指针
1. 什么是指针
当我们定义一个变量 name
```
var name string = "Go编程"
```
此时，name 是变量名，它只是编程语言中方便程序员编写和理解代码的一个标签。
当我们访问这个标签时，机算机会返回给我们它指向的内存地址里存储的值：Go编程。
出于某些需要，我们会将这个内存地址赋值给另一个变量名，通常叫做 ptr（pointer的简写），而这个变量，我们称之为指针变量。
换句话说，指针变量（一个标签）的值是指针，也就是内存地址。
根据变量指向的值，是否是内存地址，我把变量分为两种：
- 普通变量：存数据值本身
- 指针变量：存值的内存地址
2. 指针的创建
指针创建有三种方法
**第一种方法**
先定义对应的变量，再通过变量取得内存地址，创建指针
```
// 定义普通变量
aint := 1
// 定义指针变量
ptr := &aint
```
**第二种方法**
先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值。
```
// 创建指针
astr := new(string)
// 给指针赋值
*astr = "Go编程"
```
**第三种方法**
先声明一个指针变量，再从其他变量取得内存地址赋值给它
```
aint := 1
var bint *int  // 声明一个指针
bint = &aint   // 初始化
上面的三段代码中，指针的操作都离不开这两个符号：

& ：从一个普通变量中取得内存地址

*：当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量
```
通过下面这段代码，你可以熟悉这两个符号的用法
```
package main

import "fmt"

func main() {
    aint := 1     // 定义普通变量
    ptr := &aint  // 定义指针变量
    fmt.Println("普通变量存储的是：", aint)
    fmt.Println("普通变量存储的是：", *ptr)
    fmt.Println("指针变量存储的是：", &aint)
    fmt.Println("指针变量存储的是：", ptr)
}
输出如下
```
普通变量存储的是： 1
普通变量存储的是： 1
指针变量存储的是： 0xc0000100a0
指针变量存储的是： 0xc0000100a0
要想打印指针指向的内存地址，方法有两种

// 第一种
fmt.Printf("%p", ptr)

// 第二种
fmt.Println(ptr)
3. 指针的类型
我们知道字符串的类型是 string，整型是int，那么指针如何表示呢？
写段代码试验一下就知道了
```
package main

import "fmt"

func main() {
    astr := "hello"
    aint := 1
    abool := false
    arune := 'a'
    afloat := 1.2

    fmt.Printf("astr 指针类型是：%T\n", &astr)
    fmt.Printf("aint 指针类型是：%T\n", &aint)
    fmt.Printf("abool 指针类型是：%T\n", &abool)
    fmt.Printf("arune 指针类型是：%T\n", &arune)
    fmt.Printf("afloat 指针类型是：%T\n", &afloat)
}
```
输出如下，可以发现用 *+所指向变量值的数据类型，就是对应的指针类型。
```
astr 指针类型是：*string
aint 指针类型是：*int
abool 指针类型是：*bool
arune 指针类型是：*int32
afloat 指针类型是：*float64
```
所以若我们定义一个只接收指针类型的参数的函数，可以这么写
```
func mytest(ptr *int)  {
    fmt.Println(*ptr)
}
```
4. 指针的零值
当指针声明后，没有进行初始化，其零值是 nil。

func main() {
    a := 25
    var b *int  // 声明一个指针

    if b == nil {
        fmt.Println(b)
        b = &a  // 初始化：将a的内存地址给b
        fmt.Println(b)
    }
}
输出如下

<nil>
0xc0000100a0
4. 指针与切片
切片与指针一样，都是引用类型。

如果我们想通过一个函数改变一个数组的值，有两种方法

将这个数组的切片做为参数传给函数

将这个数组的指针做为参数传给函数

尽管二者都可以实现我们的目的，但是按照 Go 语言的使用习惯，建议使用第一种方法，因为第一种方法，写出来的代码会更加简洁，易读。具体你可以参数下面两种方法的代码实现

使用切片

func modify(sls []int) {
    sls[0] = 90
}

func main() {
    a := [3]int{89, 90, 91}
    modify(a[:])
    fmt.Println(a)
}
使用指针

func modify(arr *[3]int) {
    (*arr)[0] = 90
}

func main() {
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
# 第二章：面向对象
# 第三章：项目管理
# 第四章：并发编程
# 第五章：学习标准库
# 第六章：开发技能
# 第七章：第三方包
