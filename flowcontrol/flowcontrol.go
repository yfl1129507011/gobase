package flowcontrol

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Testif() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("c")
	}
}

func Testfor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Testfor2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func Testfor3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

/*
Go语言中可以使用for range 遍历数组、切片、字符串、map及通道。
通过 for range 遍历的返回值有以下规律：
1、数组、切片、字符串返回索引和值
2、map返回键和值
3、通道（channel）只返回通道内的值
*/
func Testforrange() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}
}

// switch 一个分支可以有多个值，多个case值中间使用英文逗号分隔, switch 后面需要跟变量

func Testswitch() {
	n := 1
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	case 0:
		fmt.Println("特殊")
	default:
		fmt.Println(n)
	}
}

// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
func Testswitch2() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

// goto 跳转
func Testgoto() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
breakTag:
	fmt.Println("over")
}

// 聊天机器人
func Talkrobot() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	// 读取数据直到碰到\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		// 异常退出
		os.Exit(1)
	} else {
		// 使用切片操作删除最后的\n
		name := input[:len(input)-2]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-2]

		// 全部转换为小写
		input = strings.ToLower(input)

		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}
