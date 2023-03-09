# 泛型

1. 申明一个泛型函数

``` go
package main

import "fmt"

//约定了T为any类型(any 底层就是空接口)， 形参s是元素为T类型的切片， 等价于[]interface{}
func printslice[T any](s []T){
    for _,v := range s{
        fmt.Printf("%v\n",v)
    }
}


func main(){
    printslice[int]([]int{1,2,3,4})
    printslice[float64]([]float64{1.1,2.2,3.3
    printslice[string]([]string{"aasdf","asdasdad"}

    //省略显示类型
    printslice([]int64{55,44,33,22,11})

}

```

[T any] 约束参数的类型，意思是该函数支持任何T类型;

多个泛型参数语法：

``` go
[T, M any]
[T any, M any]
[T any, M comparable]
```

2. 申明泛型切片

带有类型参数的类型被叫做泛型类型。下面定义一个底层类型为切片类型的新类型slice。它是可以存储任何类型的的切片。要使用泛型类型，要先对其进行实例化，就是给类型参数指定一个实参

``` go
// 约定参数类型为any 
type slice[T any] []T
```

``` go
package main

import "fmt"

type slice[T any] []T

func printslice[T any](s []T) {
    for _, v := range s {
        fmt.Printf("%v", v)
    }
}

func main() {
    // 初始化一个泛型类型，对应类型为int
    s1 := slice[int]{1, 2, 3}
    printslice(s1)
    // 初始化一个泛型类型，对应类型为string
    s2 := slice[string]{"123", "4r11"}
    printslice(s2)
}
```

3. 申明map 类型

``` go
// 申明一个k 为可比较类型，v 为any 类型的泛型map，在go 中k的值必须为可比较的，否则会报错 Invalid map key type: comparison operators == and != must be fully defined for the key type
type Map[K comparable,V any] map[K]V
```

``` go
package main
import "fmt"

type Map[K comparable, V any] map[K]V  

func main() {
    m1 := Map[string, int]{"key": 1}
    m1["key"] = 2

    m2 := Map[string, string]{"key": "value"}
    m2["key"] = "new value"
    fmt.Println(m1, m2)
}
```

5. 泛型约束

使用interface中规定的类型约束泛型函数的参数

``` go
package main

import "fmt"

type NumStr interface {
   Num | Str
}
type Num interface {
   ~int | ~int8 | ~int16| ~int32| ~int64| ~uint| ~ uint8| ~ uint16| ~ uint32| ~uint64| ~ uintptr| ~ float32| ~ float64| ~ complex64| ~ complex128
}
type Str interface {
   string
}

type myInt int

func add[T NumStr](a,b T) T {
   return a + b
}

func main(){
   // 等价于fmt.Println(add[int](3, 111))
   fmt.Println(add(3, 111))
   // 等价于fmt.Println(add[string]("hello", "world"))
   fmt.Println(add("hello", "world"))
   fmt.Println(myInt(1), 2) // myInt 底层也是int 类型，所以也是NumStr类型，可以使用
}
```

上面的 NumStr，新增了类型列表表达式，它是对类型参数进行约束。
使用 | 表示取并集
~int 指的是底层是int 类型的数据类型,例如自定义的myInt 类型
如果传入参数不在集合限制范围内，就会报错

使用interface中规定的方法来约束泛型的参数

```go
package main

import (
    "fmt"
    "strconv"
)
// 定义自定义类型，实现了泛型接口的方法约束
type Price int

func (i Price) String() string {
    return strconv.Itoa(int(i))
}
// 定义自定义类型，实现了泛型接口的方法约束
type Price2 string

func (i Price2) String() string {
    return string(i)
}
// 泛型接口， 有方法约束和类型约束
type ShowPrice interface {
    String() string
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
    for _, v := range s {
        ret = append(ret, v.String())
    }
    return
}

func main() {
    fmt.Println(ShowPriceList([]Price{1, 2}))
    fmt.Println(ShowPriceList([]Price2{"a", "b"}))
}
```

使用interface中规定的方法和类型来双重约束泛型的参数


```
package main

import (
    "fmt"
    "strconv"
)

// 定义自定义类型，实现了泛型接口的类型和方法约束
type Price int

func (i Price) String() string {
    return strconv.Itoa(int(i))
}

// 定义自定义类型，实现了泛型接口的类型和方法约束
type Price2 string

func (i Price2) String() string {
    return string(i)
}

// 泛型接口， 有方法约束和类型约束
type ShowPrice interface {
    String() string
    // 底层是string和int类型
    ~int | ~string
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
    for _, v := range s {
        ret = append(ret, v.String())
    }
    return
}

func main() {
    fmt.Println(ShowPriceList([]Price{1, 2}))
    fmt.Println(ShowPriceList([]Price2{"a", "b"}))
}
```

6. 自带的类型约束 any 和 comparable

any 底层就是一个空接口
any

``` go
// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
```

comparable
可比较的类型包括了booleans, numbers, strings, pointers, channels, arrays ， 所有的字段都是可比较的结构体。
该接口只是用于作为泛型约束的使用

    ```go

    // comparable is an interface that is implemented by all comparable types
// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint,
// not as the type of a variable.
type comparable interface{ comparable }
    ```

comparable也可以嵌套在自定义约束中
```go
    type ShowPrice interface {  
    int | string | comparable 
}

```

```go
package main

import (
    "fmt"
)

func findFunc[T comparable](a []T, v T) int {
    for i, e := range a {
        if e == v {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
    fmt.Println(findFunc([]string{"烤鸡", "烤鸭", "烤鱼", "烤面筋"}, "烤面筋"))
}

```

6. 注意

不能在泛型函数中申明自定义类型
例:
```go
func addIntOrFloat[T int | float64](a, b T) T {
  // 在泛型函数中声明S 的结构体
  type S struct {
      s string
  }
  s := S{"test"}
  fmt.Println(s)
  return a + b
}

func main() {
  fmt.Println(addIntOrFloat(1, 2))
}
```

结果：
```go
# generics
./main.go:41:2: type declarations inside generic functions are not   currently supported
```

. 不能直接使用泛型参数的属性
例： 声明了两个结构体Student 和Teacher 分别实现了Job的方法，在Do 泛型函数中限制参数类型为Student 和Teacher ， 并分别调用对应的job函数，此时编译报错

```go

type Student struct {
    a string
    b int
}

func (s Student) Job() {
    fmt.Println("student")
}

type Teacher struct {
    a string
    b int
}

func (t Teacher) Job() {
    fmt.Println("teaher")
}

func Do[T Student|Teacher](v T) {
    v.Job()
}
```

错误提示：

```go
generics
./main.go:26:4: v.Job undefined (type T has no field or method Job)
```

解决方案
申明一个泛型接口实现job 方法和限制的数据类型

```go
type people interface {
  *Student | *Teacher | Student | Teacher
  Job()
}
```

对应的do 函数修改如下

```go 
func Do[T people](v T) {
v.Job()
}
```