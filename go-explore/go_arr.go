package main

import "fmt"

//1 数组
func goArr() {
	//声明赋值1
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//声明赋值2
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//声明赋值3 make
	m := make([]string, 10)
	m[0] = "s"
	m[1] = "a"
	println(m)
}

//2 切片
//切片并不存储任何数据，它只是描述了底层数组中的一段。
//更改切片的元素会修改其底层数组中对应的元素
func goSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
	fmt.Println(s)

	//修改切片
	s[0] = 44
	fmt.Println(s)
}

//切片声明与取值
func goSlice2() {
	arr := [3]bool{true, true, false}  //数组初始化
	slice := []bool{true, true, false} //切片初始化：会创建一个和上面相同的数组，然后构建一个引用了它的切片
	fmt.Println(arr, slice)

	//切面取值，通过上下界操作 一个意思，取全部切片元素
	slice1 := slice[0:3]
	slice2 := slice[:3]
	slice3 := slice[0:]
	slice4 := slice[:]
	fmt.Println(slice1, slice2, slice3, slice4)

	s := slice[1:2] //true
	fmt.Println(s)

	s = slice[0:0] //[] 空
	fmt.Println(s)

	s = slice[0:1]
	fmt.Println(s)
	s = slice[2:3]
	fmt.Println(s)

	//截取
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	// 截取切片使其长度为 0
	s2 = s2[:0]
	printSlice(s2)

	// 拓展其长度
	s2 = s2[:4]
	printSlice(s2)

	// 舍弃前两个值
	s2 = s2[2:]
	printSlice(s2)

	//切片长度位零值是 nil
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s == nil {
		fmt.Println("nil!")
	}
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//make 创建切片 入参：类型，长度，容量
func goSlice3() {
	a := make([]int, 5) // len(a)=5
	fmt.Println(len(a), cap(a))
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(len(b), cap(b))
}

//append 向切片追加元素
//append 的结果是一个包含原切片所有元素加上新添加元素的切片
func goSlice4() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	//这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4, 5)
	printSlice(s)
}

//range 遍历切片或map
//for 循环的 range 形式可遍历切片或映射
func goSliceRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("循环次数为%d ，对应值为%d\n", i, v)
	}
}

//3 map映射
func goMap1() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	for i, v := range m {
		fmt.Printf("循环次数为%d ，对应值为%d\n", i, v)
	}
	//操作
	//在映射 m 中插入或修改元素：
	//m[key] = elem
	//
	//获取元素：
	//elem = m[key]
	//
	//删除元素：
	//delete(m, key)
	//
	//通过双赋值检测某个键是否存在：
	//elem, ok = m[key]
}

func main() {
	goArr()
	goSlice()
	goSlice2()
	goSlice3()
	goSlice4()
	goSliceRange()
	goMap1()
}
