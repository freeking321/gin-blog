package handler

import (
	"log"
	"net/http"
	"strings"
)

// 测试get请求和post请求参数处理
func GetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// 第一种方式
	//name := query["name"][0]
	// 第二种方式
	name := query.Get("name")
	log.Printf("get: name=%s", name)

	// 处理全部请求参数
	m := make(map[string]string)
	query2 := r.URL.Query()
	for k, v := range query2 {
		m[k] = v[0]
	}
	log.Println(m)

	// 处理路径参数  127.0.0.1:9090/get/loveyou.md
	filename := strings.Trim(r.URL.Path, "/get/")
	log.Printf(filename)
}
