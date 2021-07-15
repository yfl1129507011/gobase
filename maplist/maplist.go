package maplist

import (
	"fmt"
)

/*
map是一种无序的基于 key-value 的数据结构，Go语言中的map是引用类型，必须初始化才能使用
定义语法： map[ keyType ]ValueType
KeyType: 表示键的类型
ValueType: 表示键对应的值的类型
map类型的变量默认初始值为nil，需要使用 make() 函数来分配存储。语法为：
make( map[KeyType]ValueType, [cap] )
*/

func Testmap() {
	scoreMap := make(map[string]int, 8)
	scoreMap["job"] = 90
	scoreMap["jim"] = 100
	fmt.Println(scoreMap)                  // map[jim:100 job:90]
	fmt.Printf("type of a:%T\n", scoreMap) // type of a:map[string]int

	userinfo := map[string]string{
		"username": "fenlon",
		"password": "123456",
	}
	fmt.Println(userinfo) // map[password:123456 username:fenlon]
}

// 判断某个键是否存在
func Testmap2() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100

	// 如果key存在ok为true，v为对应的值；不存在ok为false，v为值类型零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}

// map的遍历
func Testmap3() {
	scoreMap := make(map[string]int)
	scoreMap["a"] = 12
	scoreMap["b"] = 32
	scoreMap["c"] = 42
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只想遍历key的时候
	for k := range scoreMap {
		fmt.Println(k)
	}
}

// 使用delete()函数删除键值对
func Testmap4() {
	scoreMap := make(map[string]int)
	scoreMap["a"] = 12
	scoreMap["b"] = 32
	scoreMap["c"] = 42
	delete(scoreMap, "b")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

// 元素为map类型的切片
func Testmap5() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		/*
			index:0 value:map[]
			index:1 value:map[]
			index:2 value:map[]
		*/
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	// 对切片就行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "fenlon"
	mapSlice[0]["password"] = "122222"
	mapSlice[0]["addr"] = "学院国际"
	for index, value := range mapSlice {
		/*
			index:0 value:map[addr:学院国际 name:fenlon password:122222]
			index:1 value:map[]
			index:2 value:map[]
		*/
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片类型的map
func Testmap6() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap) // map[]
	key := "中国"
	val, ok := sliceMap[key]
	if !ok {
		val = make([]string, 0, 2)
	}
	val = append(val, "a", "b")
	sliceMap[key] = val
	fmt.Println(sliceMap) // map[中国:[a b]]
}

/*
go语言中的 new 和 make 主要区别如下
1、make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据
2、new 分配返回的是指针，即类型 *Type。make 返回引用，即Type
3、new 分配的空间 被清零。make 分配空间后，会进行初始化
*/
