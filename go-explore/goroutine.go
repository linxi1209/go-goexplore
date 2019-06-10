package main

import (
	"fmt"
	"time"
)

//go程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//go 关键字：启动一个新的 Go 程并执行
func main() {
	go say("world")
	say("hello")
	doChannel()
	channelBuffer()
}

//通道
//带有类型的管道。通过 <- 来发送或者接收值；箭头就是数据流的方向
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func doChannel() {
	//对s切片求和：两通道；并行
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}

//带缓冲的通道
//make第二个参数为缓冲长度，来初始化一个带缓冲的信道
func channelBuffer()  {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}