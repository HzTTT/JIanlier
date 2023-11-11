package main

/* import(
	"fmt"
)
func main() {
	var a1 [3]int
	//初始化
	//...省略数组长度
	var a2 = [...]string{"hello","world"}
	//默认false
	//指定索引值的方式初始化
	var a3 = [3] bool{1:true}
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a3: %T\n", a3)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	for index,value := range a2{
		fmt.Printf("index: %v-----", index)
		fmt.Printf("value: %v\n", value)
	}


	//Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建
	var b1 =  new([5]int)
	//b1 的类型是 *[5]int，而 b2 的类型是 [5]int
	//////      *[]int表示一个指向整数数组的指针。
	var b2 [5]int
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b1[1]: %v\n", b1[1])
	change(b2)
	fmt.Printf("b2: %v\n", b2)
	//两种gai
	changee(&b2)
	fmt.Printf("b2: %v\n", b2)
	changee(b1)
	fmt.Printf("b1: %v\n", b1)
}
func change(a [5]int){
	a[2] = 10
}
func changee(a *[5]int){
	a[2] = 10
} */