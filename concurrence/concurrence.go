package concurrence

import (
	"fmt"
	"sync"
	"time"
)

func go1() {
	fmt.Println("go1")
}

// Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine
func Testgo1() {
	go go1()
	fmt.Println(123456)
	time.Sleep(time.Second)
}

// 启动多个goroutine
var wg sync.WaitGroup

func go2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("go2-", i)
}

func Testgo2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go go2(i)
	}
	wg.Wait()
}

// defer测试
func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func Testdefer() {
	println(f())
}
