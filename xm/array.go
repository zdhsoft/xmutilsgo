package xm

import (
	"cmp"
	"slices"
	"strings"
)

// 更多数组相关的函数，参考slices/array.go
// The algorithm based on pattern-defeating quicksort(pdqsort), but without the optimizations from BlockQuicksort.
// pdqsort paper: https://arxiv.org/pdf/2106.05123.pdf
// C++ implementation: https://github.com/orlp/pdqsort
// Rust implementation: https://docs.rs/pdqsort/latest/pdqsort/
// limit is the number of allowed bad (very unbalanced) pivots before falling back to heapsort.

// Deduplicate 去重数组元素去重
// 注意：数组元素类型必须实现了comparable接口
func Deduplicate[T comparable](paramList []T) []T {
	set := NewSetBySlice(paramList)
	return set.All()
}

// IsEqualArray 判断两个数组是否相等
// 要求数组元素顺序完全一样
// 注意：数组元素类型必须实现了comparable接口

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
// cmp比较函数 返回<0 表示按从小到大的顺序，返回>0表示从大到小的顺序
func ArraySortByFunc[S ~[]E, E interface{}](x S, cmp func(a, b E) int) {
	slices.SortFunc(x, cmp)
}

// ArraySortStableFunc 数组稳定排序（指定比较函数）
// cmp比较函数 返回<0 表示按从小到大的顺序，返回>0表示从大到小的顺序
func ArraySortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	slices.SortStableFunc(x, cmp)
}

// ArraySort 数组排序
//
//	默认的有支持整数，字符串，浮点数等可以比较大小的基本类型和衍生类型
func ArraySort[S ~[]E, E cmp.Ordered](paramArray S) {
	slices.Sort(paramArray)
}

// 有符号整数数组，连接成为字符串
//   - list 有符号整数数组
//   - sep 连接符
func IntArrayJoin[T SignedInteger](list []T, sep string) string {
	ints := make([]string, 0, len(list))
	for _, v := range list {
		ints = append(ints, I(v))
	}
	return strings.Join(ints, sep)
}

// 无符号整数数组，连接成为字符串
//   - list 无符号整数数组
//   - sep 连接符
func UIntArrayJoin[T UnsignedInteger](list []T, sep string) string {
	ints := make([]string, 0, len(list))
	for _, v := range list {
		ints = append(ints, U(v))
	}
	return strings.Join(ints, sep)
}
