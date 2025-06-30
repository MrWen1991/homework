package gorms

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBX *sqlx.DB

func init() {
	fmt.Println("main init")
	dsn := "root:123456@tcp(127.0.0.1:3306)/excersize?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = d
	fmt.Println("database initialize success....")

	// 也可以使用MustConnect连接不成功就panic
	DBX, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	DBX.SetMaxOpenConns(20)
	DBX.SetMaxIdleConns(10)
}
