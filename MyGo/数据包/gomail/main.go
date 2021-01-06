package main

import (
	"log"
	"mime"
	"strconv"
	"sync"

	"gopkg.in/gomail.v2"
)

func sendMail(mailTo string, subject, body string, wg *sync.WaitGroup) error {
	defer wg.Done()
	// mailConn126 := map[string]string{ // 使用126邮箱发送
	// 	"username": "zhangjia5551@126.com",
	// 	"authCode": "XBASUEDLZDOAPNJC", // 授权密码，而不是登陆密码
	// 	"host":     "smtp.126.com",
	// 	"port":     "465",
	// }
	mailConnQQ := map[string]string{ // 使用QQ邮箱发送
		"username": "3178945074@qq.com",
		"authCode": "ieckpqajrwlgdghf", // 授权密码，而不是登陆密码
		"host":     "smtp.qq.com",
		"port":     "465",
		// "port":     "25",  // QQ邮箱可以使用464或25端口
	}

	port, _ := strconv.Atoi(mailConnQQ["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", "Support")+"<"+mailConnQQ["username"]+">")
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConnQQ["host"], port, mailConnQQ["username"], mailConnQQ["authCode"])
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatalln("To:", mailTo, "##", "Send Email Failed!Err:", err)
	} else {
		log.Println("To:", mailTo, "##", "Send Email Successfully!")
	}
	return err
}

func main() {
	var wg sync.WaitGroup
	mailTo := []string{
		"zhangjia5551@126.com",
		"3178945074@qq.com",
	}
	subject := "Hello,Go Mail"
	body := `<html>
	<body>
		<h1 id="title">春晓</h1>
		<p class="content1">
		春眠不觉晓，
		处处闻啼鸟。
		夜来风雨声，
		花落知多少。
		</p>
	</body>
</html>`

	for _, mail := range mailTo {
		wg.Add(1)
		go sendMail(mail, subject, body, &wg)
	}
	// fmt.Println(mailConnQQ)
	wg.Wait()
}
