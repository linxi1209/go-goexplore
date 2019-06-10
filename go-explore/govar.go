package main

import (
	"fmt"
	"math/cmplx"
)

//1 直接赋值
var i, j int = 1, 2

//2 短变量声明
// 函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
func defineVar()  {
	k := 3
	fmt.Println(i, j, k)
}

//3 语法块声明并赋值
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func blockVar()  {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) //使用`%T`来输出一个值的数据类型
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//常量
//const 关键字；不能用 := 语法声明
const World = "世界"
const Truth = true

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	defineVar()
	blockVar()
}