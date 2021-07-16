package reflectfunc

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

// go反射结构体标签
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, " doc:", tagDoc)
	}
}

func Testreflect3() {
	var re resume
	findTag(&re)
	/*
		info: name  doc: 我的名字
		info: sex  doc:
	*/
}

// go 结构体标签
type movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func Testreflect4() {

	// json编码
	m := movie{"戏剧之王", 2008, 10, []string{"zhouxingchi", "zhangbozi"}}
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	// jsonStr = {"title":"戏剧之王","year":2008,"price":10,"actors":["zhouxingchi","zhangbozi"]}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// json解码
	m1 := movie{}
	err = json.Unmarshal(jsonStr, &m1)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	fmt.Printf("%v\n", m1) // {戏剧之王 2008 10 [zhouxingchi zhangbozi]}

}
