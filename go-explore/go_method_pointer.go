package main

import (
	"fmt"
	"log"
	"math"
)

type Xertex struct {
	X, Y float64
}

func (v Xertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 写法：Scale为Xertex对象的方法，入参是f
// 接收指针参数 *Xertex，常用于修改该参数
// 尝试去掉*指针，输出5，scale方法无法改变Xertex的值
func (v *Xertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	log.Println(v.X,v.Y)
}

//也可以写成
func Abs(v Xertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Xertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//注意：带指针参数的函数必须接受一个指针；
func ScaleFunc(v *Xertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
//而以指针为接收者的方法被调用时，接收者既能为值又能为指针
func (v *Xertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	log.Println(v.X,v.Y)
}
//结论：选择指针or值作为接受：
//选择指针：
//首先，方法能够修改其接收者指向的值。
//其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

func main() {
	v := Xertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	x := Xertex{3, 4}
	Scale(&x, 10)
	fmt.Println(Abs(x))

	s := Xertex{3, 4}
	v.Scale(2)
	ScaleFunc(&s, 10)
	//ScaleFunc(v, 5) 编译错误 ：带指针参数的函数必须接受一个指针


	p := &Xertex{4, 3}
	p.Scale(3) //而以指针为接收者的方法被调用时，接收者既能为值又能为指针

	ScaleFunc(p, 8)

}