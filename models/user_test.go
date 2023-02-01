package model

import (
	"fmt"
	"testing"
)

// TestMain 测试函数执行之前的函数
func TestMain(m *testing.M) {
	m.Run()
}

// TestUser 测试User 中的方法
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关的方法")

	// 通过t.Run()执行子测试函数
	t.Run("测试添加用户", testUser_AddUser) //
}

// testUser_AddUser 测试添加用户
func testUser_AddUser(t *testing.T) {
	fmt.Println("测试添加用户...")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

// testUser_De 测试删除用户
func testUser_Del(t *testing.T) {
	fmt.Println("测试删除用户...")
	user := &User{}
	user.Del()
}

// testUser_UpdateUserById 测试通过ID修改用户信息
func testUser_UpdateUserById(t *testing.T) {
	fmt.Println("测试通过ID修改用户信息...")
	user := &User{}
	user.UpdateUserById()
}

// testUser_QueryOne 测试查询一条数据
func testUser_QueryOne(t *testing.T) {
	fmt.Println("测试查询一条数据...")
	user := &User{
		ID: 15,
	}
	u, _ := user.QueryOne()
	fmt.Println("查询id为15的记录为: ", u)
}

// testUser_QueryAll 测试获取所有的数据
func testUser_QueryAll(t *testing.T) {
	fmt.Println("测试获取所有的数据...")
	user := User{}
	users, _ := user.QueryAll()
	for k, v := range users {
		fmt.Printf("第%d个用户是: %v: \n", k+1, v)
	}
}
