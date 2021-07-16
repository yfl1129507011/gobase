package interfacetype

import (
	"fmt"
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
