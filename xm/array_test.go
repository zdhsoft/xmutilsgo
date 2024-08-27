package xm

import (
	"slices"
	"testing"
)

func Test_ArraySort(t *testing.T) {
	lst := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		lst = append(lst, RandomIntScope(1, 100000))
	}
	ArraySort(lst)
	if !slices.IsSorted(lst) {
		t.Errorf("排序后的结果不是有序的！")
	}
	//  ArraySort<[]int>(lst)
}

func Test_Deduplicate(t *testing.T) {
	hasDup := func(arr []int) (bool, int) {
		seen := make(map[int]bool)
		for _, value := range arr {
			if seen[value] {
				return true, value
			}
			seen[value] = true
		}
		return false, 0
	}
	lst := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		lst = append(lst, RandomIntScope(1, 5))
	}
	newList := Deduplicate(lst)
	t.Log(lst)
	t.Log(newList)
	if has, value := hasDup(newList); has {
		t.Errorf("存在重复的元素：%d", value)
		return
	}

}
