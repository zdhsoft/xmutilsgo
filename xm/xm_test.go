package xm

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	v := Int2StringPad(1000, 20)
	if v != "00000000000000001000" {
		t.Errorf("Int2StringPad(1000, 20) != 00000000000000001000")
	}
}

func TestIntSwap(t *testing.T) {
	list := []int64{1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789, -12, -942}
	dest := []int64{1, 21, 321, 4321, 54321, 654321, 7654321, 87654321, 987654321, -21, -249}
	for idx, item := range list {
		newValue := ReverseInt64(item)
		destValue := dest[idx]
		if newValue != destValue {
			t.Errorf("ReverseInt64(%d)=%d != %d", item, newValue, destValue)
		}
	}
}

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3, 4, 5)
	s.AddFromArray([]int{1, 2, 3, 4, 5})

	s1 := NewSet(1, 2, 3, 4, 5)
	s1.AddFromArray([]int{1, 2, 3, 4, 5})

	if s.Len() != 5 {
		t.Errorf("s.Len() = %d 而不是5", s.Len())
	}

	if !s.IsEqu(s1) {
		t.Errorf("s.IsEqu(s1) != true")
	}

	testList := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: false,
		7: false,
	}

	for k, v := range testList {
		if s.Has(k) != v {
			t.Errorf("s.Has(%d) = %t ,实际要求为%t", k, s.Has(k), v)
		}
	}
	s.Remove(3, 4)

	testList[3] = false
	testList[4] = false
	for k, v := range testList {
		if s.Has(k) != v {
			t.Errorf("s.Has(%d) = %t ,实际要求为%t", k, s.Has(k), v)
		}
	}
	ss := s.All()
	fmt.Println(ss)
}

func TestIsInAndNotIn(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	notin := []int{11, 12, 13, 14, 15}
	for _, v := range s {
		if !IsIn(v, s...) {
			t.Errorf("IsIn(%d, s) != true", v)
		}
		if IsNotIn(v, s...) {
			t.Errorf("IsNotIn(%d, s) != false", v)
		}
	}
	for _, v := range notin {
		if IsIn(v, s...) {
			t.Errorf("IsIn(%d, s) != false", v)
		}
		if !IsNotIn(v, s...) {
			t.Errorf("IsNotIn(%d, s) != true", v)
		}
	}
}

func TestIsInAndNotInForArray(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	notin := []int{11, 12, 13, 14, 15}
	for _, v := range s {
		if !IsInArray(v, s) {
			t.Errorf("IsIn(%d, s) != true", v)
		}
		if IsNotInArray(v, s) {
			t.Errorf("IsNotIn(%d, s) != false", v)
		}
	}
	for _, v := range notin {
		if IsInArray(v, s) {
			t.Errorf("IsIn(%d, s) != false", v)
		}
		if !IsNotInArray(v, s) {
			t.Errorf("IsNotIn(%d, s) != true", v)
		}
	}
}

func TestEasyMD5(t *testing.T) {
	s := []string{"hello", "world", "123456", "你好，世界"}
	md5List := []string{"5d41402abc4b2a76b9719d911017c592", "7d793037a0760186574b0282f2f435e7", "e10adc3949ba59abbe56e057f20f883e", "dbefd3ada018615b35588a01e216ae6e"}

	for idx, v := range s {
		if md5List[idx] != EasyMD5(v) {
			t.Errorf("EasyMD5(%s) => %s != %s", v, EasyMD5(v), md5List[idx])
		}
	}
}

func TestPage(t *testing.T) {
	s := []*PageInfo{
		NewPage(1, 10),
		NewPage(10, 20),
		NewPage(20, 30),
		NewPage(30, 40),
		NewPage(40, 5999),
	}

	d := []*PageInfo{
		{1, 10},
		{10, 20},
		{20, 30},
		{30, 40},
		{40, 1000},
	}

	for idx, v := range s {
		if v.Page != d[idx].Page || v.PageSize != d[idx].PageSize {
			t.Errorf("s[%d] = %v, 实际要求为%v", idx, v, d[idx])
		}
	}
}

