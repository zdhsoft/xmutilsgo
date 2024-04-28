package xm

// 判断paramElement是否在paramArgs列表中
func IsIn[T comparable](paramElement T, paramArgs ...T) bool {
	for _, arg := range paramArgs {
		if paramElement == arg {
			return true
		}
	}
	return false
}

// 判断paramElement是否在paramArgs数组中
func IsInArray[T comparable](paramElement T, paramArgs []T) bool {
	for _, arg := range paramArgs {
		if paramElement == arg {
			return true
		}
	}
	return false
}

// 判断paramElement是否 不在paramArgs列表中, 不在返回true
func IsNotIn[T comparable](paramElement T, paramArgs ...T) bool {
	return !IsIn(paramElement, paramArgs...)
}

// 判断paramElement是否 不在paramArgs数组中, 不在返回true
func IsNotInArray[T comparable](paramElement T, paramArgs []T) bool {
	return !IsInArray(paramElement, paramArgs)
}
