# Chapter03 实践总结

## 这是一个简单项目的分析
优点:
- 非常工程化，把这个放在前面可以让大家了解工程结构，在学习它的同时可以浏览其他的工程
- 对工具的介绍落到了实处，每个工程都有各自的用处。
- 包管理概念提前讲解，包版本控制，包控制器等相关的讲清楚，实际的项目都是使用包，管理包，扩展包的过程。
- 命名规则，通过用例让大家代码风格趋向一致，减少阅读难度。

缺点：
- 先发展的语言，后发展的工具，可能放在前面的工具以及被替换了，或者为了适应环境做了大的更改，这个内容要大改了。
- 学习人员如果看到这个章节照着实践，很快就碰壁了，这就是周边工具影响核心功能的效果，无法验证语言的特性，无法了解写简单的代码。这就是为什么学一个语言都是从“Hello World”开始。
- 更新这个工具部分需要太多的经验，网上说明替代工具时，通常假定阅读者已经编写过一些go语言的代码，对整个流程有一定的体验。

## 工具的简单说明 核心点 go mod init指令
它的使用过程是,在targetdir 中，先写好代码 xxx.go
```shell
~user>cd targetdir #跳转到目标文件夹内
~user>go mod init com.company.yourproject.tools
~user>go mod tidy
~user>go build xxx.go
```
上面的模拟过程就是，先来到一个文件夹内，然后给这个文件夹内部的所有go文件 给它起个mod名字。
这个`com.company.yourproject.tools` 起名的方式跟unity或者java的包名字一样，一种url的逆序写法，大家按着需要写即可。
命令执行后，就有了一个固定名字的go.mod，这里面就包括了 包名，版本号，以后会继续添加需要引用的其他包的信息 用reference标记出来（go get的结果）
以下是postgresTest.go中我的修改记录跟源文件内容
```go
package main

// 这里重要介绍，没有使用的包，但是会调用它包内的公开的init函数
import (
	"database/sql"
	"fmt"
	_ "github.com/goinaction/code/chapter3/dbdriver/postgres" //这个会自动从网络下载，很强啊
)

func main() {
	db, _ := sql.Open("postgres", "mydb")
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

```

## go get
这个工具要，这个确实会download其他的包的，跟c\#的NuGet Manager很像，很好用

## go vet/go fmt
一个是检查代码，一个是格式化代码，建议提交前都使用一下。


