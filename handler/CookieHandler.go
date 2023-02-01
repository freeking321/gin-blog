package handler

import (
	"io"
	"net/http"
)

// 设置cookie 方式一
func Cookie1Handler(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "test-name",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	//设置cookie
	http.SetCookie(w, ck)

	cookie, err := r.Cookie("test-name")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, cookie.Value)
}

// 设置cookie 方式二
func Cookie2Handler(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "test-name",
		Value:  "hello2",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	//设置cookie
	w.Header().Set("set-cookie", ck.String())

	cookie, err := r.Cookie("test-name2")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, cookie.Value)
}

// CookieClear 删除Cookie
func CookieClear(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "test-name",
		Value:  "",
		Path:   "/",
		Domain: "localhost",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
