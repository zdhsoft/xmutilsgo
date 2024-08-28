package xm

import (
	"cmp"
	"slices"
)

// 更多数组相关的函数，参考slices/array.go
// The algorithm based on pattern-defeating quicksort(pdqsort), but without the optimizations from BlockQuicksort.
// pdqsort paper: https://arxiv.org/pdf/2106.05123.pdf
// C++ implementation: https://github.com/orlp/pdqsort
// Rust implementation: https://docs.rs/pdqsort/latest/pdqsort/
// limit is the number of allowed bad (very unbalanced) pivots before falling back to heapsort.

// Deduplicate 去重
func Deduplicate[T comparable](paramList []T) []T {
	set := NewSetBySlice(paramList)
	return set.All()
}

// IsEqualArray 判断两个数组是否相等
func IsEqualArray[T comparable](paramArray1, paramArray2 []T) bool {
	if len(paramArray1) != len(paramArray2) {
		return false
	}
	for i := range paramArray1 {
		if paramArray1[i] != paramArray2[i] {
			return false
		}
	}
	return true
}

// ArraySortByFunc 数组排序(指定比较函数)
func ArraySortByFunc[S ~[]E, E interface{}](x S, cmp func(a, b E) int) {
	slices.SortFunc(x, cmp)
}

// ArraySort 数组排序
//
//	默认的有支持整数，字符串，浮点数等可以比较大小的基本类型和衍生类型
func ArraySort[S ~[]E, E cmp.Ordered](paramArray S) {
	slices.Sort(paramArray)
}
