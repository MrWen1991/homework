package part4

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Secret_key = "123123"
var db *gorm.DB
var dbx *sqlx.DB

func init() {
	fmt.Println("part4 config init")
	dsn := "root:123456@tcp(127.0.0.1:3306)/weibo?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
	fmt.Println("database initialize success....")

	// 也可以使用MustConnect连接不成功就panic
	dbx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	dbx.SetMaxOpenConns(20)
	dbx.SetMaxIdleConns(10)
}
