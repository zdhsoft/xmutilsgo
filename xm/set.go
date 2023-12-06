package xm

// 集合，请用NewSetByCap或NewSet创建
type Set [T comparable] struct {
	m map[T]struct{}
}

// 增加一个或多个元素
func (s *Set[T]) Add(k ...T) {
	for _, v := range k {
		s.m[v] = struct{}{}
	}
}

// 增加一个数组的元素
func (s *Set[T]) AddFromArray(k []T) {
	for _, v := range k {
		s.m[v] = struct{}{}
	}
}

// 删除一个元素或多个元素
func (s *Set[T]) Remove(k ...T) {
	for _, v := range k {
		delete(s.m, v)
	}
}

// 删除一个数组的元素
func (s *Set[T]) RemoveFromArray(k []T) {
	for _, v := range k {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *Set[T]) Has(k T) bool {
	_, ok := s.m[k]
	return ok
}

// 清空集合
func (s *Set[T]) Clean() {
	s.m = map[T]struct{}{}
}

// 判断集合是否为空
func (s *Set[T]) IsEmpty() bool {
	return 0 == len(s.m)
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

// 指定容量创建
func NewSetByCap[T comparable](paramCaption int) Set[T] {
	m := make(map[T]struct{}, paramCaption)
	s := Set[T]{
		m: m,
	}
	return s
}

// 默认创建
func NewSet[T comparable]() Set[T] {
	m := make(map[T]struct{})
	s := Set[T]{
		m: m,
	}
	return s
}
