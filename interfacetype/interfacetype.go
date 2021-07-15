package interfacetype

import (
	"fmt"
	"reflect"
)

type animalIF interface {
	sleep()
	getType() string
	getColor() string
}

// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口
type cat struct {
	color string
}

func (c *cat) sleep() {
	fmt.Println("cat is sleep")
}

func (c *cat) getColor() string {
	return c.color
}

func (c *cat) getType() string {
	return "cat"
}

func Testinterface() {
	var animal animalIF = &cat{"red"} // 定义接口数据类型
	animal.sleep()                    // cat is sleep
	fmt.Println(animal.getColor())    // red
	fmt.Println(animal.getType())     // cat
}

type dog struct {
	color string
}

func (d *dog) sleep() {
	fmt.Println("dog is sleep")
}

func (d *dog) getColor() string {
	return d.color
}

func (d *dog) getType() string {
	return "dog"
}

func showAnimal(animal animalIF) {
	animal.sleep() // 多态
	fmt.Printf("%v\n", animal)
}

func Testinterface1() {
	cat := cat{"red"}
	dog := dog{"yellow"}
	/*
		cat is sleep
		&{red}
		dog is sleep
		&{yellow}
	*/
	showAnimal(&cat)
	showAnimal(&dog)
}

// interface{}万能数据类型
func interfaceType(arg interface{}) {
	fmt.Println(arg)

	// go语言给 interface {} 提供”类型断言“的机制
	val, ok := arg.(string)
	if !ok {
		fmt.Printf("%v is not string type\n", val)
	} else {
		fmt.Printf("%v is string type\n", val)
	}
}

func Testinterface2() {
	cat := cat{"blue"}
	interfaceType(cat)
	interfaceType(12345)
	interfaceType("yangfeilong")
}

/*
Go语言中的变量是分为两部分的:
1、类型信息：预先定义好的元信息
2、值信息：程序运行过程中可动态变化的
*/
type Reader interface {
	ReaderBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReaderBook() {
	fmt.Println("read book")
}

func (b *Book) WriteBook() {
	fmt.Println("write book")
}

func TestBook() {
	b := &Book{}

	var r Reader = b
	r.ReaderBook() // read book

	var w Write = r.(Write)
	w.WriteBook() // write book
}

// go反射 reflect
func reflectInfo(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}
func Testreflect() {
	var num float32 = 1.2324

	// type:  float32
	// value:  1.2324
	reflectInfo(num)
}

type user struct {
	Id   int
	Name string
	Age  int
}

func (u user) Call() {
	fmt.Printf("%v\n", u)
}

func getUserFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	// 通过type 获取里面的字段
	// 1、获取interface的reflect.Type，通过Type得到NumField，进行遍历
	// 2、得到每个field，数据类型
	// 3、通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v", m.Name, m.Type)
	}
}

func Testreflect2() {
	u := user{111, "dfdfd", 43}
	getUserFiledAndMethod(u)
	/*
		inputType is : user
		inputValue is: {111 dfdfd 43}
		Id:int = 111
		Name:string = dfdfd
		Age:int = 43
		Call: func(interfacetype.user)
	*/
}
