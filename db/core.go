package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义db全局变量
var Db *gorm.DB

func init() {
	// MySQL 配置信息---gorm
	var err error
	username := "root"  // 账号
	password := ""      // 密码
	host := "127.0.0.1" // 地址
	port := 3306        // 端口
	DBname := "test"    // 数据库名称
	timeout := "10s"    // 连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any("failed to connect mysql."))
	}

	// 上面那种更规范一些感觉
	//var err error
	//dsn := "root:root@tcp(127.0.0.1:3306)/gin-blog-api?charset=utf8mb4&parseTime=True&loc=Local"
	//Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true, // 禁用表名加s
	//	},
	//	Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
	//	DisableAutomaticPing:                     false,
	//	DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	//})
	//if err != nil {
	//	panic("Connecting database failed: " + err.Error())
	//}
}

// GetDB
func GetDB() *gorm.DB {
	return Db
}
