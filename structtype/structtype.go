package structtype

import "fmt"

// 类型定义和类型别名的区别
func Testtype() {
	type newInt int  // 类型定义
	type myInt = int // 类型别名

	var a newInt
	var b myInt
	// type of a:structtype.newInt  type of b:int
	fmt.Printf("type of a:%T  type of b:%T\n", a, b)
}

// 结构体的定义和实例化
func Teststruct() {
	type person struct {
		name string
		city string
		age  int8
	}

	var p1 person
	p1.name = "fenlon0"
	p1.age = 18
	p1.city = "jian"

	fmt.Printf("p1=%v\n", p1)  // p1={fenlon0 jian 18}
	fmt.Printf("p1=%#v\n", p1) // p1=structtype.person{name:"fenlon0", city:"jian", age:18}
}

// 匿名结构体
func Teststruct2() {
	var user struct {
		name string
		age  int
	}
	user.age = 25
	user.name = "dfdf"
	fmt.Printf("%#v\n", user) // struct { name string; age int }{name:"dfdf", age:25}
}

// 构造函数
type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func Teststruct3() {
	p := newPerson("a", "b", 1)
	fmt.Printf("%#v\n", p) // &structtype.person{name:"a", city:"b", age:1}
}
