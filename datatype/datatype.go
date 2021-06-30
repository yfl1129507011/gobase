package datatype

import (
	"fmt"
	"math"
)

// 整型
func Testint() {
	// 十进制
	var a int8 = 10
	fmt.Printf("%d\n", a)

	// 二进制 0b 开头
	b := 0b00101100
	fmt.Printf("%b\n", b)

	// 八进制 以0开头
	c := 077
	fmt.Printf("%o\n", c)

	// 十六进制 以0x开头
	d := 0xff
	fmt.Printf("%x\n", d)
}

// 浮点型
func Testfloat() {
	var f float32 = 1.259
	fmt.Printf("%f\n", f)

	ff := math.Pi
	fmt.Printf("%.2f\n", ff)
}

// 布尔值 类型数据只有true和false2个值
/* 注意
1、布尔型变量的默认值为false
2、go语言中不允许将整型强制转为布尔型
3、布尔型无法参与数值运算，也无法与其他类型进行转换
*/
func Testbool() {
	var isOk bool
	fmt.Println(isOk)

	isBool := true
	fmt.Println(isBool)
}

// 字符串类型  字符串的值为双引号（“）包含的内容
func Teststring() {
	s1 := "fenlon"
	// 多行字符串
	s2 := `aaa
	bbb
	ccc
	`
	// 字符串长度
	fmt.Println(s1, s2, len(s1))
}

// 类型转换  int()  float64()

/*
go语言中类型为nil的几种情况
1、空指针：var a *int
2、空列表：var a []int
*/
