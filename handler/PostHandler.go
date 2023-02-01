package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AutoGenerated struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Current  int `json:"current"`
		Size     int `json:"size"`
		Total    int `json:"total"`
		UserList []struct {
			Telephone     string `json:"telephone"`
			UserName      string `json:"userName"`
			IDCard        string `json:"idCard"`
			VipRank       string `json:"vipRank"`
			RoleName      string `json:"roleName"`
			ProjectName   string `json:"projectName"`
			AccountStatus string `json:"accountStatus"`
			CreateTime    string `json:"createTime"`
		} `json:"userList"`
	} `json:"data"`
}

// 测试post请求参数 form表单类型
func PostFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 第一种方式
	//username := r.Form["username"][0]
	//password := r.Form["password"][0]
	// 第二种方式
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Printf("username=%s, password=%s", username, password)

	// 处理全部请求参数
	m := make(map[string]string)
	form := r.PostForm
	for k, v := range form {
		if len(v) < 1 {
			continue
		}
		m[k] = v[0]
	}
	log.Println(m)
}

// 测试post请求参数 json表单类型
func PostJsonHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	log.Printf("b = %s", b)

	// 将json解析成结构体
	var postJson AutoGenerated
	json.Unmarshal([]byte(b), &postJson)
	log.Println(postJson)
}