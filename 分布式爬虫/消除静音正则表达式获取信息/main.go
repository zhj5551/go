package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// demo演示正则表达式的使用方式
func demo() {
	str := "i love zhangjia5551@126.com\nlove"
	re := regexp.MustCompile(".ve") // 完全匹配
	result := re.FindString(str)
	fmt.Println(result)
}

func main() {
	str := "https://uland.taobao.com/sem/tbsearch?spm=a2e0b.20350158.1998559106.2.4b6d468a23OVC4&refpid=mm_26632258_3504122_32538762&keyword=%E5%B9%B3%E6%9D%BF%E7%94%B5%E8%84%91&clk1=47b762da076afe6bbc69b8e8a8a7d7a0&upsId=47b762da076afe6bbc69b8e8a8a7d7a0&pid=mm_26632258_3504122_32538762&union_lens=recoveryid%3A201_11.92.48.9_5790414_1611300188006%3Bprepvid%3A201_11.92.48.9_5790414_1611300188006"

	// encoding determine for html page , eg: gbk gb2312 GB18030

	resp, err := http.Get(str)
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
		// fmt.Printf("%s\n", bodyBytes)
		parseContent(bodyBytes)
	}
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func parseContent(content []byte) {

	re := regexp.MustCompile(`<a target="_black" href="//.*>.*</a>`)
	match := re.FindAllSubmatch(content, -1)
	// fmt.Printf("%s\n", match)

	for _, v := range match {
		for _, v := range v {
			fmt.Println(string(v))
		}
	}
}
