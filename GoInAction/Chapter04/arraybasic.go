package main
import (
	"fmt"
)

func main(){ //行头不能是{标志，注意
	fmt.Println("构造数组的情况")
	//通过var的方式创建数组，并自动设置零值
	var array [5]int
	fmt.Printf("第4个元素的值是%d\n",array[4-1]) //这个输出有格式化参数，上一章的wordscount里面有
	//使用数值字面量声明数组，这里是直接assign的操作
	array2 := [5]int{10,20,30,40,50}
	fmt.Printf("第4个元素的值是%d\n",array2[4-1])
	//确定类型，让编译器自己去设置容器大小的情况
	array3 := [...]int{10,20,30,40,50}
	fmt.Printf("第4个元素的值是%d\n",array3[4-1])
	//声明数组，同时只指定部分参数的情况
	array4 := [1024]int{1:20,3:40,1023:-1}
	fmt.Printf("第4个元素的值是%d\n",array4[4-1])
	fmt.Println("指针数组，数组里面存储的是指针")
	
	
}