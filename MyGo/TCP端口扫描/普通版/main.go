package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "192.168.245.189"
	for i := 1;i<=245;i++ {
		address := fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.Dial("tcp", address)	
		if err != nil {
			fmt.Println(address, "端口关闭")
			continue
		}

		conn.Close()
		fmt.Println(address, "端口开放")
	}
}
