package common

import "testing"

//集合
//slice
func testArray(t *testing.T) {
	var slice0 [] int
	t.Log(slice0, len(slice0), cap(slice0))

	//切片填充元素
	slice0 = append(slice0, 1)
	t.Log(slice0)

	//新创建切片 int类型、初始化元素个数长度、容量
	slice1 := make([]int, 3, 5)
	t.Log(len(slice1), cap(slice1))
	//访问slice元素
	t.Log(slice1[0], slice1[2]) //长度3以内访问ok
	t.Log(slice1[3], slice1[4]) //长度3以外，尽管容量5，但会exception，因为只对slice前3位进行了初始化

	//执行填充 将len4填充上值
	slice1[3] = 2
	slice1[4] = 3
	//或直接
	slice1 = append(slice1, 2, 3)
	t.Log(slice1[3], slice1[4])
} //填充后 访问正常

func testSliceGrowing(t *testing.T) {
	s := []int{}
	//循环append 查看扩容策略
	for i := 0; i < 10; i++ {
		s = append(s, 1)
		t.Log(len(s), cap(s))
	}
}

//map
func testMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(len(m1)) //3

	m2 := make(map[int]int, 10) //初始容量
	t.Log(len(m2))              //0

	m3 := map[int]int{}
	m3[4] = 1
	t.Log(len(m3)) //1
}

//int类型value，key不存在时go也返回0，那如何判断key本身为0还是key不存在
func testKey(t *testing.T) {
	map1 := map[int]int{}
	if v, ok := map1[1]; ok { //取key时会返回一个bool，标示key存在or not
		t.Log("key exsits,v:%d", v)
	} else {
		t.Log("key doesnt exsits")
	}
}

//map的遍历
func testMapRange(t *testing.T) {
	m2 := map[int]int{1: 1, 2: 2}
	for k, v := range m2 {
		t.Log(k, v)
	}
}
