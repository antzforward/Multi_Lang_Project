package main

// 这里重要介绍，没有使用的包，但是会调用它包内的公开的init函数
import (
	"fmt"
	"database/sql"
	_"github.com/goinaction/code/chapter3/dbdriver/postgres" //这个会自动从网络下载，很强啊
)

func main(){
	db, _ :=sql.Open("postgres", "mydb") 
	defer db.Close()
	fmt.Println("Old School Work!")
}
// 【2025-4-1】 
// _"github.com/goinaction/code/chapter3/dbdriver/postgres" 看来这个方式是旧的方式。
// 给自己的文件夹起个名字，嗯 一个文件夹只要做一次就可以了
//1. go mod init mygoinaction.chapter03.mod
//2. go mod tidy
//3. go build postgresTest.go
//4. postgresTest.exe->输出 Old School Work!

//下面是我再编译另外一个文件 postgresTestFix.go
//go build postgresTestFix.go -> 提示要申请pq 于是
//1. go get github.com/lib/pq
//2. go build postgresTestFix.go
//3. postgresTestFix.exe->输出 Database connected successfully!

//以上基本上翻新了 Go in Action里面的内容。可以算小改吧，还能用。