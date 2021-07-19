package concurrence

import (
	"fmt"
	"runtime"
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

// 创建 goroutine
func TestGoroutine() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			// 退出当前整个goroutine
			runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

// go channel管道类型创建
func Testchannel() {
	// 定义channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine running...")

		c <- 123 // 将123 发送给channel管道c
	}()

	num := <-c // 从管道c接受数据，并赋值给num
	fmt.Println("num=", num)
	fmt.Println("end")
}

// 有缓冲的channel
/*
1、当channel已经满了，再向里面写数据，就会阻塞
2、当channel为空，从里面取数据也会阻塞
*/
func Testchannel2() {
	c := make(chan int, 3)

	fmt.Println("len(c)=", len(c), " cap(c)= ", cap(c))

	go func() {
		defer fmt.Println("子go程序结束")

		for i := 0; i < cap(c); i++ {
			c <- i
			fmt.Println("子go程序运行，发送的元素=", i, "len(c)=", len(c), " cap(c)= ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < cap(c); i++ {
		num := <-c
		fmt.Println("num=", num)
	}

	fmt.Println("main结束")
}

// channel 关闭
func Testchannel3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// close可以关闭一个channel
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
}

// channel 和 select
func fibo(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // c 可写，则case进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Testchannel4() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibo(c, quit)
}
