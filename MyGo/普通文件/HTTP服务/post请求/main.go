package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	requst, err := http.NewRequest("POST",
		"http://www.baidu.com",
		strings.NewReader("name=abc"))
	if err != nil {
		return
	}
	requst.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(requst)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
