package main

import (
	"fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

// 它会在程序开始运行时自动执行，且只能被声明，不能被调用。因此，您不需要（也不能）在 init 函数中添加任何实现
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	fmt.Println(prompt)
	age := 12
	//if 和 else 之后的左大括号 { 必须和关键字在同一行
	if age >= 18 {
		fmt.Println("成年了")
		//如果你使用了 else-if 结构，则前段代码块的右大括号 } 必须和 else-if 关键字在同一行。
	} else if age >= 10 {
		fmt.Println("10-18岁")
		//return
	} else {
		fmt.Println("too small")
		//return
	}

	//
	val := 10
	if val := 1; val > 5 {
		// do something
		fmt.Println("YES")
	} else {
		//if结构自己的val值，原先的val值被隐藏
		fmt.Println(val)
	}
	fmt.Println(val)
}
