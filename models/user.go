package model

import (
	"demo/utils"
	"fmt"
)

// User
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// AddUser 添加用户 - 预编译
func (user *User) AddUser() error {
	// sql语句
	sql := "insert into users(username, password, email) values (?, ?, ?)"
	// 预编译
	insertStmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("预编译出现异常: ", err)
		return err
	}
	// 执行
	_, err = insertStmt.Exec("admin3", "123456", "zhangsan@163.com")
	if err != nil {
		fmt.Println("数据写入失败: ", err)
		return err
	}
	return nil
}

// AddUser2 添加用户 Exec
func (user *User) AddUser2() error {
	// sql语句
	sql := "insert into users(username, password, email) values (?, ?, ?)"
	// 执行
	_, err := utils.Db.Exec(sql, "admin4", "123456", "lisi@163.com")
	if err != nil {
		fmt.Println("数据写入失败!: ", err)
		return err
	}

	return nil
}

// Del 根据ID删除用户
func (user *User) Del() error {
	sql := "delete from users where id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("预编译失败...")
		return err
	}

	_, err = stmt.Exec(4)
	if err != nil {
		fmt.Println("执行删除语句失败")
		return err
	}

	return nil
}

// UpdateUserById 根据ID修改用户信息
func (user *User) UpdateUserById() error {
	sql := "update users set username = ?, password = ?, email = ? where id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("预编sql失败: ", err)
		return err
	}
	_, err = stmt.Exec("admin", "123456", "123456@163.com", 4)
	if err != nil {
		fmt.Println("执行修改sql语句失败: ", err)
		return err
	}
	return nil
}

// QueryOne 查询一条记录
func (user *User) QueryOne() (*User, error) {
	// sql语句
	sql := "select * from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sql, user.ID)
	// 声明变量接受扫描出来的数据
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("查询失败:", err)
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

// QueryAll 查询全部记录
func (user *User) QueryAll() ([]*User, error) {
	sql := "select * from users"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("执行sql语句失败: ", err)
		return nil, err
	}

	// 创建User 切片
	var users []*User

	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}

		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}

		users = append(users, u)
	}
	return users, nil
}
