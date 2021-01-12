package main
import (
    "fmt"
    "net"
    "sync"
    "time"
)

func main() {

    var begin =time.Now()
    var wg sync.WaitGroup
    var ip = "192.168.245.242"
    //循环
    for j := 21; j <= 65535; j++ {
        //添加wg
        wg.Add(1)
        go func(i int) {
            //释放wg
            defer wg.Done()
            var address = fmt.Sprintf("%s:%d", ip, i)
            //conn, err := net.DialTimeout("tcp", address, time.Second*10)
            conn, err := net.Dial("tcp", address)
            if err != nil {
                fmt.Println(address, "是关闭的", err)
                return
            }
            conn.Close()
            fmt.Println(address, "打开")
        }(j)
	}
    //等待wg
    wg.Wait()
    var elapseTime = time.Now().Sub(begin)
    fmt.Println("耗时:", elapseTime)
}
