package xm

// 判断first是否在args列表中
func IsIn[T comparable](first T, args ...T) bool {
	for _, arg := range args {
		if first == arg {
			return true
		}
	}
	return false
}
func IsInArray[T comparable](first T, args []T) bool {
	for _, arg := range args {
		if first == arg {
			return true
		}
	}
	return false
}

// 判断first是否 不在args列表中
func IsNotIn[T comparable](first T, args ...T) bool {
	return !IsIn(first, args...)
}

// 判断first是否 不在args列表中
func IsNotInArray[T comparable](first T, args []T) bool {
	return !IsInArray(first, args)
}
