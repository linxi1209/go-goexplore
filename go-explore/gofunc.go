package main

import "fmt"

func main() {
	a, b := swap("hi", "max")
	fmt.Println(a, b)
	fmt.Println(split(8))
}

func swap(x, y string) (string, string) {
	return y, x
}

//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
func split(sum int) (x, y int) {
	x = sum * 2 / 4
	y = sum - x
	return //直接 返回命名的x y 变量
}
