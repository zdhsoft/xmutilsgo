package xm

import "sort"

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
func ArraySortByFunc[T any](paramArray1 []T, paramLess func(i, j int) bool) {
	sort.Slice(paramArray1, paramLess)
}

// ArraySort 数组排序
//
//	默认的有IntSlice, Float64Slice, StringSlice等类型
func ArraySort[T sort.Interface](paramArray T) {
	sort.Sort(paramArray)
}
