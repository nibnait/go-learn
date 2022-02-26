package ch1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2], arr[0])

	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1[0], arr1[1], arr1[2])

	arr2 := [...]int{1, 2, 3, 4}
	t.Log(arr2[0], arr2[1], arr2[2])
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	fmt.Println()
	for idx, elem := range arr {
		t.Log(idx, elem)
	}
	fmt.Println()
	for _, elem := range arr {
		t.Log(elem)
	}

}

//切片
func TestArraySection(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5}

	assert.Equal(t, []int{1}, arr[1:2])
	assert.Equal(t, []int{1, 2}, arr[1:3])
	assert.Equal(t, []int{1, 2, 3, 4, 5}, arr[1:])
	assert.Equal(t, []int{0, 1, 2}, arr[:3])
}
