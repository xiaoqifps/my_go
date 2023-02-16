package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mathapp/mysql"
)

type User struct {
	id   int
	username  string
	password string
}

func main() {
	if err := mysql.InitMysql(); err != nil {
		fmt.Println("数据库连接失败", err)
	}

	// 注意这个  defer 关闭的  需要  拿出来
	defer mysql.Db.Close()
	fmt.Println("数据库连接成功....")
	queryRowDemo()
	//insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
}

// 单行查询
func queryRowDemo() {
	sqlStr := "select  *  from user_tbl  where id = ?"
	var u User

	//执行查询语句, QueryRow执行完之后一定要调用  Scan 方法（会自动关闭  连接）
	row := mysql.Db.QueryRow(sqlStr, 1)
	//将数据取出赋值到  user  结构体中的变量中
	err := row.Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Println("scan  filed  fail", err)
		return
	}
	fmt.Printf("id: %d,   age: %s ,  name:%s", u.id, u.username, u.password)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user_tbl(username, password) values (?,?)"
	ret, err := mysql.Db.Exec(sqlStr, "lily", 789)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//  最后一个插入的id
	var  theID  int64
	theID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user_tbl set username=? where id = ?"
	ret, err := mysql.Db.Exec(sqlStr, "lucy", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user_tbl where id = ?"
	ret, err := mysql.Db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}






