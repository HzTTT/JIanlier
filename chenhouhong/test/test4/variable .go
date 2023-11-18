package main

/* import (
	"fmt"
	"os"
	"runtime"
)

func getInfo() (string,int){
	return "tom",15
}

func main() {
	//声明变量
	var (
		name string
		age int
		m bool
	)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("m: %v\n", m)


	//变量的类型也可以在运行时实现自动推断
	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Printf("HOME: %v\n", HOME)
	fmt.Printf("USER: %v\n", USER)
	fmt.Printf("GOROOT: %v\n", GOROOT)


	//通过 runtime 包在运行时获取所在的操作系统类型
	var goos string = runtime.GOOS
    fmt.Printf("The operating system is: %s\n", goos)
    path := os.Getenv("PATH")
    fmt.Printf("Path is %s\n", path)


	//短变量声明
	//包级别的全局变量,函数体内声明局部变量
	score :=100
	fmt.Printf("score: %v\n", score)

	
	//匿名变量
	//并行或同时赋值
	//_空白标识符（只写变量），也被用于抛弃值
	_ , year := getInfo()
	fmt.Printf("year: %v\n", year)


	//常量(编译时确定，不能接收函数返回值)
	const PI1 float64 = 3.1415
	const PI2 = 3.14
	fmt.Printf("PI1: %v\n", PI1)
	fmt.Printf("PI2: %v\n", PI2)

	//交换数值（神奇！！）
	a,b :=5,10
	a,b = b,a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	
} */