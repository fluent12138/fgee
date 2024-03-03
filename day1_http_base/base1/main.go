package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	StartLog = "==============fgee 启动==============="
)

func main() {
	fmt.Println(StartLog)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// handler为空默认走http.DefaultServeMux, 对应结构体ServeMux, 实现了Handler方法, 获取req中的Handler实例 h, 然后再执行h.ServeHTTP方法
	// 所以可以通过自定义Handler来实现自己的模块
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 打印路由
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
}

// 对应路由执行方法
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q", k, v)
	}
}
