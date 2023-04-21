package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysql_demo/mysql"
)

func main() {
	if err := mysql.InitMysql(); err != nil {
		fmt.Println("数据库连接失败", err)
	}

	// 注意这个  defer 关闭的  需要  拿出来
	defer mysql.Db.Close()
	fmt.Println("数据库连接成功....")
	mysql.QueryRowDemo()
	//mysql.InsertRowDemo()
	//mysql.UpdateRowDemo()
	//mysql.DeleteRowDemo()
}
