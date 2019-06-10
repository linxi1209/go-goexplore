package main

import (
	"fmt"
	"math"
)

//1方法
type Ver struct {
	X,Y float64
}

//特殊入参的函数 v入参 float返回值
func (v Ver) goMethod() float64 {
	return math.Abs(v.X*v.Y)
}
//等同于
func goFunc0(v Ver) float64 {
	return math.Abs(v.X*v.Y)
}

//指针接收
//入参是*T指针类型；经常用于修改传入的指针参数
func (v *Ver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//2接口
type Abser interface {
	Abs() float64
}


func main() {
	v := Ver{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v.goMethod())
}