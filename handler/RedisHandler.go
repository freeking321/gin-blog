package handler

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

var ctx = context.Background()

// StringSetValue redis sting 类型 特点 单纯的key-value  value 支持string int bool
func StringSetValue(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// 设置key-value
	err := rdb.Set(ctx, "score", 80, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	// 获取key-value 值
	val, err := rdb.Get(ctx, "score").Result()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("score 的值：%v", val)))
}

// ListSetValue redis list 类型 特点 队列操作 可以添加重复元素
func ListSetValue(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// 向balls 左侧添加 basketball
	//rdb.LPush(ctx, "balls", "basketball")
	// 向balls 右侧添加 football
	//rdb.RPush(ctx, "balls", "football")
	// 获取 balls 长度
	len, err := rdb.LLen(ctx, "balls").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)

	// 删除并返回左侧第一个元素
	//result1, err := rdb.LPop(ctx, "balls").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("val1: ", result1)
	// 删除并返回右侧第一个元素
	//result2, err := rdb.RPop(ctx, "balls").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("val2: ", result2)

	// 返回list队列 区间内的元素
	strings, err := rdb.LRange(ctx, "balls", 0, 4).Result()
	for k, v := range strings {
		fmt.Printf("k: %v, v: %v \n", k, v)
	}

	w.Write([]byte(fmt.Sprintf("balls 的长度：%v", len)))
}

// SetSetValue redis list 类型 特点 集合 可以求交集 差集 不可以添加重复元素
func SetSetValue(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// 向 AFriends 增加三个好友
	rdb.SAdd(ctx, "AFriends", "mary")
	rdb.SAdd(ctx, "AFriends", "job")
	rdb.SAdd(ctx, "AFriends", "bitch")

	// 向 BFriends 增加三个好友
	rdb.SAdd(ctx, "BFriends", "job")
	rdb.SAdd(ctx, "BFriends", "dem")

	// 判断元素是否在集合中
	ok, err := rdb.SIsMember(ctx, "AFriends", "bitch").Result()
	fmt.Println(ok)
	if err != nil {
		log.Fatal(err)
	}

	// 取交集
	r2, _ := rdb.SInter(ctx, "AFriends", "BFriends").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交集的好友是: ", r2)

	// 取差集  a没在b 中的都取出来
	r3, _ := rdb.SDiff(ctx, "AFriends", "BFriends").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("差集的好友是: ", r3)

	// 返回所有元素
	r4, _ := rdb.SMembers(ctx, "BFriends").Result()
	fmt.Println("全部元素: ", r4)
}

func HashSetValue(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		// PoolSize: 20, // 通过PoolSize设置初始化连接池数量 redis连接池
	})

	// 向key为 c_user_info 中添加 name 和 age
	rdb.HSet(ctx, "c_user_info", "name", "杨柳")
	rdb.HSet(ctx, "c_user_info", "age", 20)

	// 返回key c_user_info 中name 和 age 的值
	r1, _ := rdb.HGet(ctx, "c_user_info", "name").Result()
	fmt.Println("name: ", r1)
	r2, _ := rdb.HGet(ctx, "c_user_info", "age").Result()
	fmt.Println("age: ", r2)

	// 批量添加  批量获取  删除
}
