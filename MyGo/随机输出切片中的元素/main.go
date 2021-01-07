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
    }
    str := slice[rand.Intn(len(slice))]
    fmt.Println(str)
}

