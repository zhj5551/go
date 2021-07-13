
## http的继承与路由控制
```
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port string
)


type muxHandler struct {
	handlers    map[string]http.Handler
	handleFuncs map[string]func(resp http.ResponseWriter, req *http.Request)
}

// 初始化多路由操作器
func newMuxHandler() *muxHandler {
	return &muxHandler{
		make(map[string]http.Handler),
		make(map[string]func(resp http.ResponseWriter, req *http.Request)),
	}
}

func (handler *muxHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	if hl, ok := handler.handlers[urlPath]; ok {
		hl.ServeHTTP(resp, req)
		return
	}
	if fn, ok := handler.handleFuncs[urlPath]; ok {
		fn(resp, req)
		return
	}
	http.NotFound(resp, req)
}

func (handler *muxHandler) handle(pattern string, hl http.Handler) {
	handler.handlers[pattern] = hl
}
func (handler *muxHandler) HandleFunc(pattern string, fn func(resp http.ResponseWriter, req *http.Request)) {
	handler.handleFuncs[pattern] = fn
}



type baseHander struct {
	result string
}

func (hander *baseHander) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path=>", req.URL.Path)
	fmt.Println("url param a =>", req.URL.Query().Get("a"))
	resp.Write([]byte(hander.result))
}

func main() {
  // 参数解析
	flag.StringVar(&port, "port", ":8080", "port to listen")
  参数解析
	flag.Parse()
	router := newMuxHandler()
	// router.
	a := baseHander{}
	a.result = "zhangjia"

	b := baseHander{}
	b.result = "bbbb"
	// a.ServeHTTP()
	router.handle("/hello/a/", &a)
	router.handle("/hello/b/", &b)
	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})
  // 使用router代替默认的server
	http.ListenAndServe(port, router)

}

```
