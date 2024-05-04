package xm

// 集合，请用NewSetByCap或NewSet创建
type Set[T comparable] struct {
	m map[T]struct{}
}

// 增加一个或多个元素
func (s *Set[T]) Add(paramElement ...T) {
	for _, v := range paramElement {
		s.m[v] = struct{}{}
	}
}

// 增加一个数组的元素
func (s *Set[T]) AddFromArray(paramElementArray []T) {
	for _, v := range paramElementArray {
		s.m[v] = struct{}{}
	}
}

// 删除一个元素或多个元素
func (s *Set[T]) Remove(paramElement ...T) {
	for _, v := range paramElement {
		delete(s.m, v)
	}
}

// 删除一个数组的元素
func (s *Set[T]) RemoveFromArray(paramElementArray []T) {
	for _, v := range paramElementArray {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *Set[T]) Has(paramElement T) bool {
	_, ok := s.m[paramElement]
	return ok
}

// 清空集合
func (s *Set[T]) Clean() {
	s.m = map[T]struct{}{}
}

// 判断集合是否为空
func (s *Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

// 取集合的数量
func (s *Set[T]) Len() int {
	return len(s.m)
}

// 取集合所有元素
func (s *Set[T]) All() []T {
	list := []T{}
	for i := range s.m {
		list = append(list, i)
	}
	return list
}

// 取交集
func (s *Set[T]) And(paramOther Set[T]) Set[T] {
	needLen := s.Len()
	olen := paramOther.Len()
	if olen > needLen {
		needLen = olen
	}
	resultSet := NewSetByCap[T](needLen)

	for i := range s.m {
		if paramOther.Has(i) {
			resultSet.Add(i)
		}
	}
	return resultSet
}

// 取并集
func (s *Set[T]) Or(paramOther Set[T]) Set[T] {
	needLen := s.Len() + paramOther.Len()
	resultSet := NewSetByCap[T](needLen)

	for i := range s.m {
		resultSet.Add(i)
	}

	for i := range paramOther.m {
		resultSet.Add(i)
	}
	return resultSet
}

// 取差集
func (s *Set[T]) Diff(paramOther Set[T]) Set[T] {
	resultSet := NewSetByCap[T](s.Len() + paramOther.Len())
	for i := range s.m {
		if !paramOther.Has(i) {
			resultSet.Add(i)
		}
	}
	for i := range paramOther.m {
		if !s.Has(i) {
			resultSet.Add(i)
		}
	}
	return resultSet
}

// 其他集合元素没有在本集合中的
func (s *Set[T]) NotInBySet(paramOther Set[T]) []T {
	list := make([]T, 0, paramOther.Len())
	for i := range paramOther.m {
		if !s.Has(i) {
			list = append(list, i)
		}
	}
	return list
}

// 其他数组元素没有在本集合中的
func (s *Set[T]) NotInByArray(paramOther []T) []T {
	OtherSet := NewSetBySlice(paramOther)
	list := make([]T, 0, len(paramOther))
	for i := range s.m {
		if !OtherSet.Has(i) {
			list = append(list, i)
		}
	}
	return list
}

// 其他数组元素没有在本集合中的
func (s *Set[T]) InByArray(paramOther []T) []T {
	list := []T{}
	if paramOther == nil {
		return list
	}
	list = make([]T, 0, len(paramOther))
	for _, i := range paramOther {
		if s.Has(i) {
			list = append(list, i)
		}
	}
	return list
}

// 判断两个集合是否相等
func (s *Set[T]) IsEqu(paramOther Set[T]) bool {
	if s.Len() != paramOther.Len() {
		return false
	}
	for i := range s.m {
		if !paramOther.Has(i) {
			return false
		}
	}
	for i := range paramOther.m {
		if !s.Has(i) {
			return false
		}
	}
	return true
}

// 指定容量创建
func NewSetByCap[T comparable](paramCaption int) Set[T] {
	m := make(map[T]struct{}, paramCaption)
	s := Set[T]{
		m: m,
	}
	return s
}

// 默认创建
func NewSet[T comparable](paramList ...T) Set[T] {
	m := make(map[T]struct{})
	for _, v := range paramList {
		m[v] = struct{}{}
	}
	s := Set[T]{
		m: m,
	}
	return s
}

// 通过切片创建集合
func NewSetBySlice[T comparable](paramSlice []T) Set[T] {
	m := make(map[T]struct{}, len(paramSlice))
	for _, v := range paramSlice {
		m[v] = struct{}{}
	}
	s := Set[T]{
		m: m,
	}
	return s
}
