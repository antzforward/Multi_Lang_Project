package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // 空白导入触发驱动注册
)

func main() {
	// 标准连接字符串格式
	dsn := "user=postgres password=secret dbname=testdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected successfully!")
}
