package main

import "fmt"

//指针
//指针保存了值的内存地址
//* 表示指针指向的底层值 返回的是指针指向的值
//& 会生成一个指向其操作数的指针 返回的是指针
func goPointer() {
	i, j := 2, 3
	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针修改 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

//结构体
func goStruct() {
	v := Vertex{1, 2}
	v.X = 4
	println(v.X, v.Y)
}

type Vertex struct {
	X int
	Y int
}

//结构体指针
func goStructPointer() {
	v := Vertex{1, 2}
	p := &v //如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以go允许我们使用隐式间接引用，直接写 p.X 就可以。
	p.X = 33
	fmt.Println(v)
}

func goStructOpt() {
	v1 := Vertex{1, 2} // 创建一个 Vertex 类型的结构体
	v2 := Vertex{X: 1} // Y:0 被隐式地赋予
	v3 := Vertex{}     // X:0 Y:0
	p := &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	fmt.Println(v1, v2, v3, p)

}

func main() {
	goPointer()
	goStruct()
	goStructPointer()
	goStructOpt()
}
