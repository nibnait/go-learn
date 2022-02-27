package ch1

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log("len m1", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2", len(m2))

	m3 := make(map[int]int, 10)
	t.Log("len m3", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	m2 := map[int]string{}
	t.Log(m2[1])

	m2[2] = ""
	t.Log(m2[2])

	m1[3] = 3
	if value, isExist := m1[3]; isExist {
		t.Logf("key 3 存在。 value = %d", value)
	} else {
		t.Log("key 3 不存在")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
