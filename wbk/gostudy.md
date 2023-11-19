# 1.学习资料

进阶项目：https://dev.to/techschoolguru/design-db-schema-and-generate-sql-code-with-dbdiagram-io-4ko5

基础学习：https://www.bookstack.cn/read/the-way-to-go_ZH_CN/eBook-directory.md

git指南：https://www.liaoxuefeng.com/wiki/896043488029600

# 2.go环境配置

（windows配置一键完成，没啥奇怪的地方，pass）

（linux等需要时再去配吧，先跳过了）

# 3.go常用命令与问题

## 1. go mod init （name）

cd一个文件夹，在终端中执行改命令，创建一个go.mod，name写啥无所谓，跟文件夹名不要求一致，但一致可能会方便查阅

这玩意目前来看，是类似java的package的东西，貌似不认文件夹名

## 2.go的main函数

整个go的项目中只能有一个main，与其他语言不同的一点是，它要求写一个main包，类似以下代码

```go
package main

import(
	"fmt"
)

func main(){

}
```

**不要求go文件名为main**

（ps：这也算是一个代码模版吧）

## 3.go run （name）.go

要执行代码，需要在终端写以上命令

## 4.go test

貌似是用来测试用例的，没用过，到时候再来个测评吧

## 5.关于工作区

当你打开了一个最大的文件夹后，再在里面分别写两个单独文件夹的go项目（各自一个main包，各自一个mod）会爆红，应该是vscode的识别问题，解决方法，只打开该go项目的文件夹即可，不要打开最大的文件夹。

## 6.go install

导入外部安装包

```go
go install codesite.ext/author/goExample/goex
```



# 4.常量与变量

（一丢丢语法区别）

```go
//常量
const Pi = 3.14159
//枚举（有一点特别）
const (
    Unknown = 0
    Female = 1
    Male = 2
)
// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
    a = iota  // a = 0
    b         // b = 1
    c         // c = 2
    d = 5     // d = 5   
    e         // e = 5
)
```

```go
//变量
var a, b *int
//也可以
var (
    a int
    b bool
    str string
)

//自动识别类型
(变量名):= (值)
```

# 5.基本类型

**布尔类型** 

bool

**整数：**
int
uint
uintptr
int8（-128 -> 127）
int16（-32768 -> 32767）
int32（-2,147,483,648 -> 2,147,483,647）
int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
uint8（0 -> 255）
uint16（0 -> 65,535）
uint32（0 -> 4,294,967,295）
uint64（0 -> 18,446,744,073,709,551,615）

**浮点型（IEEE-754 标准）：**
float32（+- 1e-45 -> +- 3.4 * 1e38）
float64（+- 5 1e-324 -> 107 1e308） >>>>一般建议用它

**复数类型(%v)：**
complex64 (32 位实数和虚数)
complex128 (64 位实数和虚数)

```go
var c1 complex64 = 5 + 10i

//实部和虚部都为float
c = complex(re, im)

//函数 real(c) 和 imag(c) 可以分别获得相应的实数和虚数部分
```





## 类型转换

```go
//将变量whole强转为int
int(whole)
```

## 字符串

可以用+拼接字符串

# 6.控制结构

## 1.if-else结构

```go
if condition1 {
    // do something    
} else if condition2 {
    // do something else    
} else {
    // catch-all or default
}

//比较有意思的，initialization只是一个初始化语句，不参与判断
if initialization; condition {
    // do something
}
```

## 2.测试多返回值函数的错误

举例一下吧，目前还没用上

```go
value, err := pack1.Function1(param1)
if err != nil {
    fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
    return err
}
// 未发生错误，继续执行：

//习惯用法
if err != nil {
    fmt.Printf("Program stopping with error %v", err)
    os.Exit(1)
}
```

## 3.switch

```go
//不需要break就会自动跳出，从上往下搜索答案

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

```go
//希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
switch i {
    case 0: fallthrough
    case 1:
        f() // 当 i == 0 时函数也会被调用
}
```

## 4.for

```go
//语法
for 初始化语句; 条件语句; 修饰语句 {}

//举例
    for i := 0; i < 5; i++ {
        fmt.Printf("This is the %d iteration\n", i)
    }
```

## 5.标签与goto

就是那个能到处跳的玩意，嗯，要求标签一定要在goto之前，一般也不会用，pass

没啥特别的，也是语法区别，上例子

```go
package main
import "fmt"
func main() {
LABEL1:
    for i := 0; i <= 5; i++ {
        for j := 0; j <= 5; j++ {
            if j == 4 {
                continue LABEL1
            }
            fmt.Printf("i is: %d, and j is: %d\n", i, j)
        }
    }
}
```

```go
package main
func main() {
    i:=0
    HERE:
        print(i)
        i++
        if i==5 {
            return
        }
        goto HERE
}
```

# 7.函数

## 1.go特色

go的函数支持多返回值，但是return不能跟着一个表达式，这是错的（顺路提一嘴：i++这种在go里不能作为表达式赋值，即：不可以n = i++）

## 2.语法

```go
//举例更好懂
func Test(a, b int) (x1, x2, x3 int) {
	x1 = a + b
	x2 = a * b
	x3 = a - b
	return x1, x2, x3
}

//或者
func getX2AndX3(input int) (int, int) {
    return 2 * input, 3 * input
}
```

## 3.传递变长参数

啥是变长参数，简单来说，就是你不知道要传多少个参数，就设置了个类似于数组的参数在那

```go
//使用函数部分
    slice := []int{7,9,3,5,1}
    x = min(slice...)
//也可以直接传多个参数
 x := min(1, 3, 2, 0)

//函数定义部分
func min(s ...int) int {
    if len(s)==0 {
        return 0
    }
    min := s[0]
    for _, v := range s {
        if v < min {
            min = v
        }
    }
    return min
}
```

## 4.defer和追踪

defer就是延后这段代码的执行（等全部执行完了再执行，如果有多个defer，执行顺序参考入栈出栈）

> 举例

```go
func f() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}

//输出
4 3 2 1 0
```

## 5.内置函数（直接查文档就知道了）

## 6.将函数作为参数

还行，就语法问题

```go
//使用
callback(1, Add)

//函数
func Add(a, b int) {
    fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
    f(y, 2) // this becomes Add(1, 2)
}
```

## 7.闭包

不取函数名，类似于java的匿名内部类

```go
func() {
    sum := 0
    for i := 1; i <= 1e6; i++ {
        sum += i
    }
}()

//不取函数名，且最后的()是表示对该匿名函数的调用，等于手动传参（？）
```



有意思的是，函数可以赋给变量

```go
    for i := 0; i < 4; i++ {
        g := func(i int) { fmt.Printf("%d ", i) }
        g(i)
        fmt.Printf(" - g is of type %T and has value %v\n", g, g)
    }
```

### 1.闭包应用：函数作为返回值

虽然还不知道有什么大用处，先记一下吧

```go
func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}
```

### 2.闭包调试

不太能理解怎么应用

```go
var where = log.Print
func func1() {
where()
... some code
where()
... some code
where()
}
```

## 8.计算函数执行时间

用函数time.Now()即可

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

# 9.数组

## 1、语法

```go
var identifier [len]type

//举例
var arr1 [5]int
```

## 2.特殊点

```go
var arrLazy = [...]int{5, 6, 7, 8, 22}

//部分赋值
var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
```

# 10.切片

声明切片的格式是： `var identifier []type`（不需要说明长度）。



切片的初始化格式是：`var slice1 []type = arr1[start:end]`。

这表示 `slice1` 是由数组 `arr1` 从 `start` 索引到 `end-1` 索引之间的元素构成的子集（切分数组，`start:end` 被称为切片表达式）。

> 修改切片长度

slice1 = slice1[:len(slice1)-1]

## 1.将切片传递给函数

```go
func sum(a []int) int {
    s := 0
    for i := 0; i < len(a); i++ {
        s += a[i]
    }
    return s
}
func main() {
    var arr = [5]int{0, 1, 2, 3, 4}
    sum(arr[:])
}
```

## 2.用 make() 创建一个切片

> 用法

```go
make([]int, 50, 100)
new([100]int)[0:50]
```

# 11.For-range 结构

第一个返回值 `ix` 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 `for` 循环内部可见的局部变量。`value` 只是 `slice1` 某个索引位置的值的一个拷贝，不能用来修改 `slice1` 该索引位置的值。

```go
for ix, value := range slice1 {
    ...
}
```

# 12.map

用过都说好（doge）

```go
var map1 map[keytype]valuetype
var map1 map[string]int
```

```
v := map1[key1]` 可以将 `key1` 对应的值赋值给 `v
```

`len(map1)` 方法可以获得 `map` 中的 pair 数目

## 1.测试键值对是否存在

```go
_, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false
```

## 2.删除

从 `map1` 中删除 `key1`：

直接 `delete(map1, key1)` 就可以。

## 3.for-range 的配套用法

```go
for key, value := range map1 {
    ...
}
```

## 4.map的切片（多看）

```go
 items := make([]map[int]int, 5)
    for i:= range items {
        items[i] = make(map[int]int, 1)
        items[i][1] = 2
    }
```

# 13.包

## 1.regexp包

```go
searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
pat := "[0-9]+.[0-9]+" //正则
ok, _ := regexp.Match(pat, []byte(searchIn))
ok, _ := regexp.MatchString(pat, searchIn)
```

感觉还是整个copy下来，看得通透点

```go
package main
import (
    "fmt"
    "regexp"
    "strconv"
)
func main() {
    //目标字符串
    searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
    pat := "[0-9]+.[0-9]+" //正则
    f := func(s string) string {
        v, _ := strconv.ParseFloat(s, 32)
        return strconv.FormatFloat(v*2, 'f', 2, 32)
    }
    if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
        fmt.Println("Match Found!")
    }
    re, _ := regexp.Compile(pat)
    //将匹配到的部分替换为"##.#"
    str := re.ReplaceAllString(searchIn, "##.#")
    fmt.Println(str)
    //参数为函数时
    str2 := re.ReplaceAllStringFunc(searchIn, f)
    fmt.Println(str2)
}
```

## 2.自定义包

当写自己包的时候，要使用短小的不含有 `_`（下划线）的小写单词来为文件命名



主程序利用的包必须在主程序编写之前被编译。主程序中每个 pack1 项目都要通过包名来使用：`pack1.Item`

所以如果想不写这个前缀包名

<hr>

```go
import . "./pack1"
```

当使用 `.` 作为包的别名时，你可以不通过包名来使用其中的项目。例如：`test := ReturnStr()`。

## 3.godoc工具

差不多是提供一个文档给自己的团队看

在命令行中执行

```go
godoc -http=:6060 -goroot="."

//（. 是指当前目录，-goroot 参数可以是 /path/to/my/package1 这样的形式指出 package1 在你源码中的位置或接受用冒号形式分隔的路径，无根目录的路径为相对于当前目录的相对路径）
```

然后去浏览器中打开地址

```go
http://localhost:6060
```

就可以看到godoc页面啦

## 4.go install

就是拉取安装一些外部包



更新到新版本的 Go 之后本地安装包的二进制文件将全被删除。如果你想更新，重编译、重安装所有的 go 安装包可以使用：`go install -a`。

# 14.结构 (struct) 与方法 (method)

## 1.结构体定义

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

## 2.**选择器 (selector)**

其实就是常用的那个成员指示符  .

## 3.方法和函数

> 函数

func 函数名（传入参数） 传出参数 即可

> 方法

func (结构体实例) 函数名 （传入参数） 传出参数

结构体实例：这里称为 接收者，接收者可以是指针，也可以不是，但是实际用的时候不受使用者到底是否指针的影响：

```go
package main
import (
    "fmt"
)
type B struct {
    thing int
}
func (b *B) change() { b.thing = 1 }
func (b B) write() string { return fmt.Sprint(b) }
func main() {
    var b1 B // b1 是值
    b1.change()
    fmt.Println(b1.write())
    b2 := new(B) // b2 是指针
    b2.change()
    fmt.Println(b2.write())
}
/* 输出：
{1}
{1}
*/
```



这个就很符合java的类的成员方法了，限定只有这个结构体的实例可用

举例：

```go
type TwoInts struct {
    a int
    b int
}

two1 := new(TwoInts)
 two1.AddThem()


func (tn *TwoInts) AddThem() int {
    return tn.a + tn.b
}
```

注：已有的（比如什么int啥的）不能直接加方法上去的，但是可以间接加：**先定义该类型（比如：`int` 或 `float32(64)`）的别名类型，然后再为别名类型定义方法**

