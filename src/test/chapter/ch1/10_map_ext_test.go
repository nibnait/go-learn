package ch1

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	judgeExist(t, 1, mySet)

	// 添加元素
	mySet[2] = true
	judgeExist(t, 2, mySet)

	// 删除元素
	delete(mySet, 2)
	judgeExist(t, 2, mySet)

}

func judgeExist(t *testing.T, i int, set map[int]bool) {
	if set[i] {
		t.Logf("%d 存在", i)
	} else {
		t.Logf("%d 不存在", i)
	}
}
