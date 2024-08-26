package xm

import (
	"sort"
	"testing"
)

func Test_ArraySort(t *testing.T) {
	lst := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		lst = append(lst, RandomIntScope(1, 100000))
	}
	t.Log("----", lst)
	sort.Ints(lst)
	t.Log(">>>>", lst)
	//  ArraySort<[]int>(lst)
}
