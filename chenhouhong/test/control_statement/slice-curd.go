package main

/* import(
	"fmt"
)
func main() {
	a := make([]int, 10)
	fmt.Printf("a: %v\n", a)
	//添加一个在末尾
	a = append(a, 3)
	fmt.Printf("a[10]: %v\n", a[10])

	slFrom := []int{1, 2, 3}
    slTo := make([]int, 10)
	//slTo 新数组
    n := copy(slTo, slFrom)
    fmt.Println(slTo)
    fmt.Printf("Copied %d elements\n", n) // n == 3
    sl3 := []int{1, 2, 3}
	//func append(s[]T, x ...T)
	//append在超过容量时会分配新切片，此时切片指向不同的相关数组
	//切片相连append(a13,,slFrom...)
    sl3 = append(sl3, 4, 5, 6,11,1,1,1,1,1,1,1,1,1)
    fmt.Println(sl3)


	s := "你好"
	by :=[]byte(s)
	//dst :=[]byte		错误
	var dst []byte
	copy(dst ,s)
	// r := []rune(s)		//元素类型
	// fmt.Printf("r: %v\n", r)

	//必须s...????
	by = append(by, s...)
	fmt.Printf("by: %v\n", by)
	for index,value := range by {
		fmt.Printf("index: %v\n", index)
		fmt.Printf("value: %c\n", value)
	}
	//for-range 字符串可以输出中文
	for index,value := range s {
		fmt.Printf("index: %v\n", index)
		fmt.Printf("value: %c\n", value)
	}
	str :=s[0:3]	//获取子字符串
	fmt.Printf("str: %v\n", str)
} */
