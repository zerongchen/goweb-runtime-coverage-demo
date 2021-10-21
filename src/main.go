package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWord(s string) {
	fmt.Println("say:  ", s)
}

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // 解析参数，默认是不会解析的
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	handleRadon(r.Form["number"][0], w)
}

func main() {
	http.HandleFunc("/", test)               // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
