package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GrammarHandler() {
	str := "word "
	sl := []byte(str)
	fmt.Println(string(sl[1]))

	// 1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}

	//3.timer实现延时的功能
	//(1)
	time.Sleep(time.Second)
	//(2)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	//(3)
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")

	//4.停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}

	//5.重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
}

//--------------------------语法学习分割线------------------------------------
//func init() {
//	lens := len("1234s")
//	fmt.Println(lens)
//}
//
//func init() {
//	lens := len("中国")
//	fmt.Println(lens)
//}
//

//func main() {
//	go func() {
//		defer fmt.Println("A.defer")
//		func() {
//			defer fmt.Println("B.defer")
//			// 结束协程
//			runtime.Goexit()
//			defer fmt.Println("C.defer")
//			fmt.Println("B")
//		}()
//		fmt.Println("A")
//	}()
//	for {
//	}
//}

//func main() {
//	//r := gin.Default()
//	//r.GET("/ping", func(c *gin.Context) {
//	//	c.JSON(200, gin.H{
//	//		"message": "pong",
//	//	})
//	//})
//	//// get query 参数
//	//r.GET("/user", func(c *gin.Context) {
//	//	name := c.DefaultQuery("name", "枯藤")
//	//	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
//	//})
//	//// post 表单 参数
//	//r.POST("/form", func(c *gin.Context) {
//	//	types := c.DefaultPostForm("type", "post")
//	//	username := c.PostForm("username")
//	//	password := c.PostForm("userpassword")
//	//	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
//	//})
//	//// 路由组
//	//v1 := r.Group("/v1")
//	//{
//	//	v1.GET("/login", login)
//	//	v1.GET("/submit", submit)
//	//}
//	//
//	//v2 := r.Group("/v1")
//	//{
//	//	v2.POST("/login", login)
//	//	v2.POST("/submit", submit)
//	//}
//	//
//	//// 404页面
//	//r.NoRoute(func(c *gin.Context) {
//	//	c.String(http.StatusNotFound, "404 not found2222")
//	//})
//	//
//	//r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
