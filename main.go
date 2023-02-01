package main

import (
	"demo/handler"
	"log"
	"net/http"
)

func main() {
	// 测试get和post
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/post/form", handler.PostFormHandler)
	http.HandleFunc("/post/json", handler.PostJsonHandler)
	// 测试cookie
	http.HandleFunc("/cookie1", handler.Cookie1Handler)
	http.HandleFunc("/cookie2", handler.Cookie2Handler)
	// 测试redis
	http.HandleFunc("/redis/string", handler.StringSetValue)
	http.HandleFunc("/redis/list", handler.ListSetValue)
	http.HandleFunc("/redis/set", handler.SetSetValue)
	http.HandleFunc("/redis/hash", handler.HashSetValue)

	log.Println("Running at port 8899 ...")

	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
