package xm

// 集合，请用NewSetByCap或NewSet创建
type Set[T comparable] struct {
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
func (s *Set[T]) And(other *Set[T]) Set[T] {
	if other == nil {
		return *s
	}

	needLen := s.Len()
	olen := other.Len()
	if olen > needLen {
		needLen = olen
	}
	resultSet := NewSetByCap[T](needLen)

	for i := range s.m {
		if other.Has(i) {
			resultSet.Add(i)
		}
	}
	return resultSet
}

// 取并集
func (s *Set[T]) Or(other *Set[T]) Set[T] {
	if other == nil {
		return *s
	}

	needLen := s.Len() + other.Len()
	resultSet := NewSetByCap[T](needLen)

	for i := range s.m {
		resultSet.Add(i)
	}

	for i := range other.m {
		resultSet.Add(i)
	}
	return resultSet
}

// 取差集
func (s *Set[T]) Diff(other *Set[T]) Set[T] {
	if other == nil {
		return *s
	}
	resultSet := NewSetByCap[T](s.Len() + other.Len())
	for i := range s.m {
		if !other.Has(i) {
			resultSet.Add(i)
		}
	}
	for i := range other.m {
		if !s.Has(i) {
			resultSet.Add(i)
		}
	}
	return resultSet
}

// 其他集合元素没有在本集合中的
func (s *Set[T]) NotInBySet(other *Set[T]) []T {
	list := []T{}
	if other == nil {
		return list
	}
	list = make([]T, 0, other.Len())
	for i := range other.m {
		if !s.Has(i) {
			list = append(list, i)
		}
	}
	return list
}

// 其他数组元素没有在本集合中的
func (s *Set[T]) NotInByArray(other *[]T) []T {
	list := []T{}
	if other == nil {
		return list
	}
	list = make([]T, 0, len(*other))
	for _, i := range *other {
		if !s.Has(i) {
			list = append(list, i)
		}
	}
	return list
}

// 其他数组元素没有在本集合中的
func (s *Set[T]) InByArray(other *[]T) []T {
	list := []T{}
	if other == nil {
		return list
	}
	list = make([]T, 0, len(*other))
	for _, i := range *other {
		if s.Has(i) {
			list = append(list, i)
		}
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
