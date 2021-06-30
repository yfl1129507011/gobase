package function

import "fmt"

/*
go语言中定义函数使用 func 关键字，具体格式如下:
func 函数名(参数)(返回值) { 函数体 }
- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。
- 参数：由参数变量和参数变量类型组成
- 返回值：由返回值变量和返回值变量类型组成，多个返回值必须使用 （）包裹，并用，分隔
*/

func IntSum(x int, y int) int {
	return x + y
}

// 可变参数 指函数的参数数量不固定。go语言中的可变参数通过在参数名后加 ... 来标识
func IntSumAll(x ...int) int {
	fmt.Println(x) // x是一个切片
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func Intcalc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

/*
变量作用域
1、全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效，在函数中可以直接访问全局变量
2、函数内定义的变量无法在该函数外使用；如果局部变量和全局变量重名，优先访问局部变量
*/
var Num int32 = 120

func Testnum() {
	Num := 100
	fmt.Println(Num) // 100
}

// 匿名函数
func Testniming() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	// 自执行函数：匿名函数定义完加（）直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

// 闭包 指一个函数和与其相关的引用环境组合而成的实体。闭包=函数(返回值为函数)+引用环境
func add() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func Testadd() {
	f := add()
	fmt.Println(f(1)) // 1
	fmt.Println(f(1)) // 2
	fmt.Println(f(1)) // 3
	fmt.Println(f(1)) // 4
}

// defer 语句：go语言中 defer 语句会将其后面跟随的语句进行延迟处理。
// 在 defer 归属的函数即将放回是，将延迟处理的语句按 defer 定义的逆序进行执行
func Testdefer() {
	fmt.Println("a")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("b")

	/*
		输出结果如下：
		a
		b
		3
		2
		1
	*/
}
