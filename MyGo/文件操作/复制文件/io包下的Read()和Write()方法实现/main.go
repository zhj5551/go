package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile1 我们可以通过io包下的Read()和Write()方法，边读边写，就能够实现文件的复制。这个方法是按块读取文件，块的大小也会影响到程序的性能。
// 该函数的功能：实现文件的拷贝，返回值是拷贝的总数量(字节),错误
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。。")
			break
		} else if err != nil {
			fmt.Println("报错了。。。")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}

func main() {
	n, err := CopyFile1(`C:\Users\zhj5551\go\src\go\MyGo\aa.html`, "aa.html")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(n)
}
