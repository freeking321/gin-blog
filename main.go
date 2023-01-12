package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
//	str := "word "
//	sl := []byte(str)
//	fmt.Println(string(sl[1]))
//
//	//// 1.timer基本使用
//	//timer1 := time.NewTimer(2 * time.Second)
//	//t1 := time.Now()
//	//fmt.Printf("t1:%v\n", t1)
//	//t2 := <-timer1.C
//	//fmt.Printf("t2:%v\n", t2)
//	//
//	//// 2.验证timer只能响应1次
//	//timer2 := time.NewTimer(time.Second)
//	//for {
//	//	<-timer2.C
//	//	fmt.Println("时间到")
//	//}
//
//	// 3.timer实现延时的功能
//	//(1)
//	//time.Sleep(time.Second)
//	//(2)
//	//timer3 := time.NewTimer(2 * time.Second)
//	//<-timer3.C
//	//fmt.Println("2秒到")
//	//(3)
//	//<-time.After(2*time.Second)
//	//fmt.Println("2秒到")
//
//	// 4.停止定时器
//	//timer4 := time.NewTimer(2 * time.Second)
//	//go func() {
//	// <-timer4.C
//	// fmt.Println("定时器执行了")
//	//}()
//	//b := timer4.Stop()
//	//if b {
//	// fmt.Println("timer4已经关闭")
//	//}
//
//	// 5.重置定时器
//	//timer5 := time.NewTimer(3 * time.Second)
//	//timer5.Reset(1 * time.Second)
//	//fmt.Println(time.Now())
//	//fmt.Println(<-timer5.C)
//}
