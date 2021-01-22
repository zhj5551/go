package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// url := "https://www.douban.com/"
	// url := "https://www.baidu.com/"
	// utf8编码
	// url := "https://uland.taobao.com/sem/tbsearch?spm=a2e0b.20350158.1998559106.2.4b6d468a23OVC4&refpid=mm_26632258_3504122_32538762&keyword=%E5%B9%B3%E6%9D%BF%E7%94%B5%E8%84%91&clk1=47b762da076afe6bbc69b8e8a8a7d7a0&upsId=47b762da076afe6bbc69b8e8a8a7d7a0&pid=mm_26632258_3504122_32538762&union_lens=recoveryid%3A201_11.92.48.9_5790414_1611300188006%3Bprepvid%3A201_11.92.48.9_5790414_1611300188006"
	url := "http://www.chinanews.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("err status code : %d\n", resp.StatusCode)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(result))
}
