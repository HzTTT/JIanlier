package main

/* import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	//返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败
	//也可以使用一个 error 类型的变量来代替作为第二个返回值：成功执行的话，error 的值为 nil，否则就会包含相应的错误信息（Go 语言中的错误类型为 error: var err error
	//if 语句来测试执行结果；由于其符号的原因，这样的形式又称之为“逗号 ok 模式”(comma, ok pattern)。
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize) 
	str := "hello world"
	num,err := strconv.Atoi(str)
	if err != nil{
		//习惯用法
		//fmt.Printf("An error occured in pack1.Function1 with parameter %v\n", str)

		//但当输出到文件流、网络流等具有不确定因素的输出对象时，应该始终检查是否有错误发生
		count,er := fmt.Println(err)
		fmt.Printf("count: %v\n", count)
		fmt.Printf("er: %v\n", er)		//nil或0-----true,error-----输出到文件流、网络流等具有不确定因素的输出对象时
		//（此处的退出代码 1 可以使用外部脚本获取到）
		os.Exit(1)
		//return 
	}else{
		fmt.Println(num)
	}




} */