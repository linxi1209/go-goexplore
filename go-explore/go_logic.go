package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

//1 循环
//Go 只有一种循环结构：for 循环
func goFor() {
	sum := 0
	for i := 0; i < 10; i++ { //初始化语句；条件表达式；后置语句
		sum += i
	}
	fmt.Println(sum)
}

//初始化语句和后置语句是可选
func goFor2() {
	sum := 1
	for ; sum < 10; {
		sum += sum
	}
	fmt.Println(sum)
}

//for 是 Go 中的 “while”
func goForWhile() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

}

//无限循环
//省略循环条件，该循环就不会结束
func goForever() {
	for {
		fmt.Println("hi")

	}
}

//if
// 表达式外无需小括号 ( ) ，而大括号 { } 则是必须的
func goIf1(x float64) string {
	if x < 0 {
		return goIf1(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//if 的简短语句
func goIf2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //直接跟上结果判断  pow：求x的n次方（次幂）
		return v
	}
	return lim

}

//if else
//在 if 的简短语句中声明的变量，同样可以在任何对应的 else 块中使用；
func goIf3(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了；v的声明周期只在 if else语块中
	return lim
}

//switch
//Go 自动提供了每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。
//Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数
func goSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//switch 的 case 语句从上到下顺次执行，直到匹配成功时停止
func goSwitch2() {
	today := time.Now().Weekday()
	switch time.Monday {
	case today:
		println("today.")
	case today + 1:
		println("tomorrow.")
	default:
		println("Too far away.")

	}
}

//没有条件的 switch 同 switch true 一样
func goSwitch3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("morning")
	case t.Hour() < 17:
		println("afternoon")
	default:
		println("evening")

	}
}

//defer
//在函数return之后，延迟执行
func goDefer()  {
	defer fmt.Println("world")
	fmt.Print("hello ")
}

//defer栈 n个defer会被压入一个栈中。当外层函数return后，defer函数会按照后进先出的顺序调用
func goDefer2()  {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	goFor()
	goFor2()
	goForWhile()
	//goForever()

	fmt.Println(goIf1(2), goIf1(-4))
	fmt.Println( //块 写法
		goIf2(3, 2, 10),
		goIf2(3, 3, 20),
	)

	goSwitch()
	goSwitch2()
	goSwitch3()

	goDefer()
	goDefer2()
}