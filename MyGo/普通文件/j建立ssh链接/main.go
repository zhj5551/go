package main

import (
	"bytes"
	"fmt"
	"log"

	// crypto 安装方法：进到$GOPATH/golang.org/x 目录，运行git clone https://github.com/golang/crypto.git
	"golang.org/x/crypto/ssh"
)

func main() {
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", "192.168.80.131:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("5")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("yun install -y nginx"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
