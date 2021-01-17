package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "zs", "姓名")
	age := flag.Int("age", 18, "年龄")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
