package variable

import "fmt"

/*
变量注意事项
1、函数外的每个语句都必须以关键字开始（var、const、func等）
2、:= 不能使用在函数外
3、_ 多用于占位，表示忽略值
*/

// 变量标准声明以关键字var开头
var name string
var age int
var isOk bool

// 变量批量声明
var (
	s string
	i int
	b bool
	f float32
)

/*
变量的初始化
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。
每个变量会被初始化成其类型的默认值，例如：整形和浮点型的默认值为0；
字符串变量的默认值为空字符串；布尔型的变量默认值为false；
切片、函数、指针变量的默认值nil
*/
var aName string = "fenlon"
var aAge int = 25

// 类型推导初始化变量
var bName, bAge = "yfl", 18

// 在函数内部可以使用 := 方式声明并初始化变量
func test_1() {
	cName := "abc"
	cAge := 12
	fmt.Println(cName, cAge)
	fmt.Printf("cName的数据类型是：%T\n cAge的数据类型是：%T\n", cName, cAge)
}

// 匿名变量 使用下划线 _ 表示 (匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复申明)
func foo() (int, string) {
	return 100, "ageLimit"
}

func testFoo() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x, "y=", y)
}

/*
常量  使用关键字 const 进行申明
*/
const pi = 3.1415

const (
	limit = 1000
	e     = 2.7182
)

// const同时申明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2 // 100
	n3 // 100
)

// iota 在const关键字出现时将被重置为0。
// const中每新增一行常量申明将使 iota 计数一次 (iota可理解为const语句块中的行索引)
const (
	day1 = iota // 0
	day2        // 1
	day3        // 2
	day4        // 3
	day5        // 4
)

// 在iota中使用 _ 可以跳过某些值
const (
	floor1 = iota // 0
	floor2        // 1
	floor3        // 2
	_             // 3
	floor4        // 4
)

// iota 申明中间插队
const (
	room1  = iota // 0
	room2  = 100  // 100
	room3  = 0    // 0
	rooom4 = iota // 4
	room5         // 5
)
const room6 = iota // 0

// 多个 iota 定义在一行
const (
	id1, id2 = iota + 1, iota + 2 // 1, 2
	id3, id4                      // 3, 4
	id5, id6                      // 5, 6
)
