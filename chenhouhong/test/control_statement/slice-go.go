package main

/* import(
	"fmt"
)
func main() {
	//切片（slice),底层增加自动扩容，拥有相同数据类型的可变长度序列，优点：是引用类型
	var k = [3]int{1,2,3}
	fmt.Printf("k: %v\n", k)
	//1）定义切片
	var a1 []int
	fmt.Printf("a1: %v\n", a1)
	//数组未初始化之前默认位nil
	fmt.Print(a1 == nil)

	//2)切片是引用数据类型，用make函数创建切片
	var a2 = make([]int,3)
	fmt.Printf("\na2: %v\n", a2)
	//需要先扩展切片的大小，append函数
	 //a2[10] = 100	?????
	fmt.Printf("len(a2): %v\n", len(a2))
	fmt.Printf("cap(a2): %v\n", cap(a2))



	//切片的初始化
	//数组的复制，普通数组也可以[1:3]
	//使用数组初始化
	var b1 = []int{1,2,3,4,5,6}		//注：初始化得到的实际上是切片slice
	b2 := b1[1:3]		//[)包左不包右
	fmt.Printf("b2: %v\n", b2)
	b2[1] = 100
	//影响b1后面应该b3
	b3 := b1[:]
	fmt.Printf("b3: %v\n", b3)


	c1 := [...]int{1,2,3,4}
	fmt.Printf("c1: %v\n", c1)
	c2 := c1[1:2]
	fmt.Printf("c2: %v\n", c2)

	for index,value := range c1{
		fmt.Printf("index: %v\n", index)
		fmt.Printf("value: %v\n", value)
	}


	j := [5]int{1,2,3,4,5}
	var jj []int = j[2:5]
	fmt.Printf("jj: %v\n", jj)
	jj[0] = 99
	//相关数组也改变了
	fmt.Printf("j[2]: %v\n", j[2])


	slice := make([]int,50,100)
	slice[0] = 19
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("len(slice): %v\n", len(slice))
	fmt.Printf("cap(slice): %v\n", cap(slice))

	//字符串是纯粹不可变的字节数组，它们也可以被切分成切片

	for index,_ := range slice {
		slice[index] = 1
		fmt.Printf("v: %v\n", slice[index])
	}
	
} */