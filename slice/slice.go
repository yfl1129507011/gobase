package slice

import "fmt"

/*
数组的长度是固定的并且长度属于类型的一部分
切片是一个拥有相同类型元素的可变长度的序列
*/

func Testslice() {
	var a []string              // 声明一个字符串切片
	var b = []int{}             // 声明一个整型切片并初始化
	var c = []bool{false, true} // 声明一个布尔切片并初始化

	fmt.Println(a) // []
	fmt.Println(b) // []
	fmt.Println(c) // [false true]

	/*
		切片不能直接比较, 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	*/
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
}

/*
切片表达式
1、简单切片表达式( a[low:high] )：表达式中的low和high表示一个索引范围（左包含，右不包含即 low<=索引值<high）
省略了 low 则默认为0; 省略了 high 则默认为切片操作数的长度
2、完整切片表达式( a[low : high : max] ) ：简单切片表达式a[low: high]相同类型、相同长度和元素的切片。
另外，它会将得到的结果切片的容量设置为max-low。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0
*/

// 简单切片表达式
func Testslice2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)) // s:[2 3] len(s):2 cap(s):4
}

// 完整切片表达式
func Testslice3() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t)) // t:[2 3] len(t):2 cap(t):4
}

/*
使用 make() 函数构造切片，格式如下：
make( []T, size, cap ) , 其中：
T：切片的元素类型
size：切片中元素的数量
cap：切片的容量
*/
func Testslice4() {
	a := make([]int, 2, 10)
	fmt.Println(a)      // [0 0]
	fmt.Println(len(a)) // 2
	fmt.Println(cap(a)) // 10
}

// 切片的赋值拷贝
func Testslice5() {
	s1 := make([]int, 3)
	s2 := s1
	s2[0] = 100
	fmt.Println(s1) // [100 0 0]
	fmt.Println(s2) // [100 0 0]
}

// 切片遍历
func Testslice6() {
	s := []int{1, 3, 5}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for key, val := range s {
		fmt.Println(key, val)
	}
}

// append()方法为切片添加元素
func Testslice7() {
	var s []int
	s = append(s, 1) // [1]
	fmt.Println(s)
	s = append(s, 2, 3, 4) // [1 2 3 4]
	fmt.Println(s)
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]
	fmt.Println(s)
}

// append() 添加元素和切片策略性扩容
func Testslice8() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

// 使用 copy() 函数复制切片
func Testslice9() {
	a := []int{1, 2, 3}
	c := make([]int, 3)
	copy(c, a)
	fmt.Println(a) // [1 2 3]
	fmt.Println(c) // [1 2 3]

	c[0] = 1000
	fmt.Println(a) // [1 2 3]
	fmt.Println(c) // [1000 2 3]
}
