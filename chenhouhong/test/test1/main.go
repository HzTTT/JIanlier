package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
 "test/test2"
 
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入一个字符串: ")
	input, _ := reader.ReadString('\n')
	output := strings.ToUpper(input)
	fmt.Println("转换后的字符串:", output)
	s:=test2.Hello()
	fmt.Printf("s:%v\n",s)
}