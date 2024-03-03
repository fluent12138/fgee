package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUNT %s\n", req.URL)
	}
}

func main() {
	// 自定义一个处理请求的结构体
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
