package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// 这是用的go的原生包 go-sql-driver
	Db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("mysql connect failed err: ", err.Error())
		panic(any(err.Error()))
	}
}
