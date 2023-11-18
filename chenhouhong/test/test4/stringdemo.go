package main

/* import (
	//"bytes"
	"fmt"
	"strconv"
	//"strings"
)

func main() {
// 	//字节定长数组，无法修改
// 	a := "hello"
// 	//反引号，\n原样输出
// 	//模板字符串？
// 	b :=`hello
// nihao
// \n
// world`
// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Print(string(a[1])+"\n")
// 	//网络上表单，拼URL……
// 	//字符串拼接
// 	message := fmt.Sprintf("a = %s,b = %s",a,b)
// 	fmt.Printf("message: %v\n", message)

// 	fmt.Println("=======")
// 	//字符串拼接
// 	str := `hell`+`o`
// 	str+=",world"
// 	fmt.Printf("str: %v\n", str)
// 	s := strings.Join([]string{a, b}, "、")
// 	fmt.Printf("s: %v\n", s)
// 	//缓冲(效率更高)
// 	var buffer bytes.Buffer
// 	buffer.WriteString("how ")
// 	buffer.WriteString("old ")
// 	buffer.WriteString("are ")
// 	buffer.WriteString("you")
// 	fmt.Printf("buffer.String(): %v\n", buffer.String())
// 	fmt.Println("=======")

// 	//字符串的切片操作
// 	strs := "hello world"
// 	n := 3
// 	m := 6
// 	fmt.Println(strs[n:])
// 	fmt.Println(strs[:m])	//包左不包右
// 	fmt.Println(strs[n:m])	//[n,m)
// 	fmt.Println(len(strs))



	// //字符串常见方法
	// str := " Hello World llo "
	// fmt.Printf("len(str): %v\n", len(str))
	// fmt.Printf("strings.Split(str, \" \"): %v\n", strings.Split(str, " "))
	// fmt.Printf("strings.Contains(str, \"lo\"): %v\n", strings.Contains(str, "lo"))
	// fmt.Printf("strings.HasPrefix(str, \"hello\"): %v\n", strings.HasPrefix(str, "hello"))
	// fmt.Printf("strings.HasSuffix(str, \"World\"): %v\n", strings.HasSuffix(str, "World"))
	// fmt.Printf("strings.Index(str, \"llo\"): %v\n", strings.Index(str, "llo"))
	// fmt.Printf("strings.LastIndex(str, \"llo\"): %v\n", strings.LastIndex(str, "llo"))
	// fmt.Printf("strings.ToLower(str): %v\n", strings.ToLower(str))
	// fmt.Printf("strings.Replace(str, \"llo\", \"kko\", 2): %v\n", strings.Replace(str, "llo", "kko", 2))
	// fmt.Printf("strings.Count(str, \"llo\"): %v\n", strings.Count(str, "llo"))
	// ss := strings.Repeat(str,3)
	// fmt.Printf("ss: %v\n", ss)
	// fmt.Printf("strings.Trim(str, \"llo\"): %v\n", strings.Trim(str, "llo"))
	// fmt.Printf("strings.TrimSpace(str): %v\n", strings.TrimSpace(str))


	// str := " hello world "
	// //返回字符串数组,or 返回slice？？for-range 循环
	// ss := strings.Fields(str)
	// fmt.Printf("ss: %v\n", ss[1])
	// ss = strings.Split(str ," ")
	// fmt.Printf("ss: %v\n", ss[1])
	// //有间隔符的拼接
	// ss1 :=strings.Join([]string{str,"world","wide"}," kkk ")
	// fmt.Printf("ss1: %v\n", ss1)
	// //strings.NewReader(str)??????
	

	a := 15016
	//不是iota
	var str string =strconv.Itoa(a)
	fmt.Printf("str: %v\n", str)
	A,err := strconv.Atoi(str)
	fmt.Printf("b: %v\n", A)
	fmt.Printf("err: %v\n", err)
	//bitSize参数，因为是转换成float64，所以无影响？
	B,err1 :=strconv.ParseFloat("0.26545121",5)
	fmt.Printf("B: %v\n", B)
	fmt.Printf("err1: %v\n", err1)

} */