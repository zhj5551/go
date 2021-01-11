package main
import (
	"fmt"
	"sync"
	"time"
	"net"
)

func main() {
	begin := time.Now()
	var wg sync.WaitGroup
	ip := "192.168.245.189"
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ip, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Println(address,"端口关闭")
				return
			}
			conn.Close()
			fmt.Println(address, "打开")
		}(port)	
	}	
	wg.Wait()
	elapseTime := time.Now().Sub(begin)
	fmt.Println("运行时间：", elapseTime)					
}
