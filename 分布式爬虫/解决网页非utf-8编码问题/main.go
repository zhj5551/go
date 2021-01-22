package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// encoding determine for html page , eg: gbk gb2312 GB18030
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	e := determineEncoding(resp.Body)
	reader := transform.NewReader(resp.Body, e.NewDecoder())
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", bodyBytes)
	}
}
