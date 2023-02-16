package mysql

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func InitMysql() (err error) {

	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db"
	//Open  函数只是校验   dsn  的查数是否正确，  并不会连接数据库
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("连接成功？？？")

	//尝试与数据库进行连接
	err = Db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	return
}