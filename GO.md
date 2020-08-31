# GO 语言学习

## 一、基础

### 1.1 包，变量和函数

#### 1.1.1包 package

每一个Go程序都是由package构成的。程序从main包开始运行

``` go
package main

import "fmt"
/*import(
	"fmt"
	"math"
)*/
func main(){
    fmt.Println("hello world")
}
```

按照约定，包名与导入路径的最后一个元素一致。例如，`"math/rand"` 包中的源码均以 `package rand` 语句开始。

#### 1.1.2 导入

```go
import "fmt"
import "math"
```

#### 1.1.3 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza` 就是个已导出名，`Pi` 也同样，它导出自 `math` 包。

#### 1.1.4 函数

```go
package main

import "fmt"

func add(x int, y int) int {//最后面的int是返回值类型
	return x + y
}
/*
func add(x, y int) int { 
	return x + y
}
*/

func main() {
	fmt.Println(add(42, 13))
}

```

函数可以莫得参数或者很多参数

#### 1.1.5 变量

`var` 语句用于声明一个变量列表，类型在最后

```go
var c, python, java bool
```

变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
    i := 1 //这种声明不能在函数外部使用
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```



#### 1.1.6 Go语言的基本类型

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```



表达式 `T(v)` 将值 `v` 转换为类型 `T`。

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```



#### 1.1.7 常量

常量的声明与变量类似，只不过是使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明。







### 1.2for if else switch 和 defer

#### 1.2.1 for

Go只有一种循环结构：`for` 循环

基本的 `for` 循环由三部分组成，它们用分号隔开：

- 初始化语句：在第一次迭代前执行（optional）
- 条件表达式：在每次迭代前求值
- 后置语句：在每次迭代的结尾执行（optional）

初始化语句通常为一句短变量声明，该变量声明仅在 `for` 语句的作用域中可见。

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```

for 是 Go中的 “while”

#### 1.2.2 if

Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的。

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

else 就类似于c语言的else

#### 1.2.3 switch 

十分的类似于C，C++的switch。但是GO自动给每个case后面提供了`break` 并且不是每一个case都要执行，可以通过 `fallthrough` 来结束语句。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {//这句话表示对os进行switch，并且在还在之前给os这个string变量进行了一个赋值
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

switch的case语句是从上到下的

也存在没有条件的switch语句，就类似于 `switch true` 

```go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

#### 1.2.4 defer

defer语句回将函数推迟到外层函数返回之后执行，推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

推迟的函数会被压入一个stack，类似于栈的工作方式

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

这个输出结果是 done 9 8 7 6 5 4 3 2 1 0





### 1.3strut slice 和映射

#### 1.3.1 指针

```go
var p *int
```

这个p就是一个指针

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。

`&` 操作符会生成一个指向其操作数的指针。

```
i := 42
p = &i
```



#### 1.3.2 结构体

一个结构体（`struct`）就是一组字段（field）。这个看上去和C语言的结构体没啥区别，对于结构体字段用点号来访问。

如果我们有一个指向结构体的指针 `p`，那么可以通过 `(*p).X` 来访问其字段 `X`。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 `p.X` 就可以。

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, *p, v2, v3)
}
```





#### 1.3.3 数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

表达式

```
var a [10]int
```

会将变量 `a` 声明为拥有 10 个整数的数组。





#### 1.3.4 切片 slide

切片则为数组元素提供动态大小的，灵活的视角

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```
a[low : high]
```

low表示第一个元素，high表示最后一个元素 **但是切片不包括high** 

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

这里发现s是对应primes下表1，2，3这三个数

切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

**关于切片的长度和容量** len() 就是切片包含的元素的个数，cap()表示切片从它的第一个元素开始，到底层数组元素末尾的个数。

切片的零值是 `nil`

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

**用make创建切片**

```go
a := make ([]int,5)//len(a)=5,cap(a)=5
a := make ([]int,0,5]) //len(a)=0,cap(a)=5
```

**向切片追加元素** Go语言提供了内建的append 函数

```go
func append(s []T, vs ...T) []T
```

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

output:

```markdown
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```



#### 1.3.5 Range

`for` 循环的 `range` 形式可遍历切片或映射。

当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

可以将下标或值赋予 `_` 来忽略它。

```
for i, _ := range pow
for _, value := range pow
```

若你只需要索引，忽略第二个变量即可。

```
for i := range pow
```



#### 1.3.6 映射 这是不是C++的map或者set呀

映射将键映射到值

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

```

这个vertex就可以被省略

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```















## 二、方法和接口

### 2.1 方法

#### 2.1.1 方法

go没有类，但是可以给struct定义方法。

```go
func(variable_name variable_data_type) function_name() [return_type]{
	/*函数本体*/
}
```

之前学c++我们都知道有类函数，并且类函数还有public和private。但是Go语言里面没有类，怎么给结构体定义函数呢

```go
package main

import "fmt"

/* 定义结构体 */
type Circle struct {
  radius float64 //定义了一个结构体
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {// c是变量，Circle是变量类型，这里是结构体。 getArea是方法名称，float64是方法返回值类型
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```

我们要知道Go语言并不是一个面向对象语言。



#### 2.1.2

当然上面的代码也可以写成

```go
package main

import "fmt"

type Circle struct {
  radius float64 //定义了一个结构体
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", getArea(c1))
}

//该 method 属于 Circle 类型对象中的方法
func getArea(c Circle) float64{
  return 3.14 * c.radius * c.radius
}
```

只不过这样子就是一个标准的函数，上面的方法只有Circle这个结构体才能用这个函数

#### 2.1.3 指针接收者

这个看的不是很懂但是代码还是能看懂理解的

```go
func(variable_name variable_data_type) function_name() [return_type]{
	/*函数本体*/
}
```

这个基本规则中的variable_data_type 可以是一个指针。

并且function_name()这个括号里面也可以添加参数

eg：

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

这个程序的运行结果就是v的x点和y点都被扩大了十倍，间距也扩大了10倍。





### 2.2 接口

#### 2.2.1 接口类型

接口类型是由一组方法签名定义的集合

```go
type interface_name interface{
	method_name 1 [return_type]
	method_name 2 [return_type]
	method_name 3 [return_type]
	...
}//定义接口

type struct_name struct {
    /*variables*/
}//定义结构体

func(struct_name_variable struct_name) method_name1() [return_type]{
    /*implement*/
}
...
```

这是一个interface的基本定义

关于接口的隐式实现我没看懂



#### 2.2.2 Stringer

fmt包中定义的Stringer是最普遍的接口之一







## 三、并发*



#### 3.1 Go程

goroutine 是一个线程

```go
go f(x,y,z)
```

会启动一个新的Go程序并且执行

#### 3.2 信道 channel

信道是带有类型的管道。

可以通过神奇的信道操作符号 <- 来发送或者接受值。

```go
ch <- v //将v发送至信道ch
v := <-ch //从ch接收值并赋予v
//信道和映射切片一样，在使用前必须被建立
ch := make (chan int)
```

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	y, x := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
```

