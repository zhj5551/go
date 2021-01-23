package main

// 把素数写到管道中
func addNum(num int, addChan chan int) {
	for i := 1; i <= num; i++ {
		addChan <- i
	}
	close(addChan)
}

func readData(addChan chan int, sushuChan chan int) {
	for {
		num := <-addChan
		// flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				break
			}
			sushuChan <- num
		}

	}
}

func main() {

}
