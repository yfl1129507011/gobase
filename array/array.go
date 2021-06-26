package array

import "fmt"

/*
数组的声明语法： var 数组变量名 [元素数量]Type类型
*/
func Testarray() {
	var testArr [3]int        // 数组会初始化为int类型的零值
	var numArr = [3]int{1, 2} //使用指定的初始值
	var cityArr = [3]string{"北京", "上海", "广州"}

	fmt.Println(testArr) // [0 0 0]
	fmt.Println(numArr)  // [1 2 0]
	fmt.Println(cityArr) // [北京 上海 广州]
}

// 初始值的个数自行推断数组的长度
func Testarray2() {
	var testArr [3]int
	var numArr = []int{1, 2}
	var cityArr = []string{"北京", "上海", "广州"}

	fmt.Println(testArr)                        // [0 0 0]
	fmt.Println(numArr)                         // [1 2 0]
	fmt.Printf("type of numArr:%T\n", numArr)   // type of numArr:[]int
	fmt.Println(cityArr)                        // [北京 上海 广州]
	fmt.Printf("type of cityArr:%T\n", cityArr) // type of cityArr:[]string
}

// 指定索引值的方式来初始化数组
func Testarray3() {
	a := []int{1: 3, 5: 7}
	fmt.Println(a) // [0 3 0 0 0 7]
}

// 数组遍历
func Testarrayview() {
	a := []int{1: 3, 5: 7}

	// for 循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// for range遍历
	for k, v := range a {
		fmt.Println(k, v)
	}
}

// 二维数组
func Testtwoarray() {
	a := [][2]string{
		{"曹操", "曹丕"},
		{"刘备", "刘禅"},
		{"孙权", "孙皓"},
	}

	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
