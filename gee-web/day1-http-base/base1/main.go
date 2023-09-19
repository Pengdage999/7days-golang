package main

// 创建了一个 HTTP 服务器
// 监听于本地的端口 9999
// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HandleFunc注册一个处理器函数handler和对应的模式pattern。
	http.HandleFunc("/", indexHandler)      //注册了一个处理根路径 "/" 的处理函数
	http.HandleFunc("/hello", helloHandler) //注册了一个处理路径 "/hello" 的处理函数
	// 启动了一个 HTTP 服务器，监听在本地的 9999 端口
	// 如果出现错误，它会调用 log.Fatal() 来记录并退出程序
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
