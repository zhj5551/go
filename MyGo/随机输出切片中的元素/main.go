package main

import (
    "fmt"
    "math/rand"
    "time"
)


func main() {
    rand.Seed(time.Now().UnixNano())
    slice := []string{
        "aa bb cc",
        "bbbbb",
        "ccccc",
        "ddddd",
        "eee",
        "ff",
		"hh",
    }
	for {
    	str := slice[rand.Intn(len(slice))]
    	fmt.Println(str)
		time.Sleep(time.Second / 2)
	}
}

