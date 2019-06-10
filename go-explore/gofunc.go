package main

import (
	"fmt"
	"math"
)


func swap(x, y string) (string, string) {
	return y, x
}

//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
func split(sum int) (x, y int) {
	x = sum * 2 / 4
	y = sum - x
	return //直接 返回命名的x y 变量
}

//函数值
//函数也是值。它们可以像其它值一样传递 用作参数或返回值
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func goFunc() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	//调用compute func hypot作为参数传入
	fmt.Println(compute(hypot))
}

//函数的闭包 todo 不是太理解；应用场景？
//Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
//
//例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Println("xxx",x)
		sum += x
		return sum
	}
}
func goPack() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//斐波纳契闭包
// 返回一个“返回int的函数”
func fibonacci() func() int {
	//todo 。。。
	return nil
}

func doFibonacci()  {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func main() {
	//a, b := swap("hi", "max")
	//fmt.Println(a, b)
	//fmt.Println(split(8))

	goFunc()
	goPack()
}