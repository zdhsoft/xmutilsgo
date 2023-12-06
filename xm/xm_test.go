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

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3, 4, 5)
	s.AddFromArray([]int{1, 2, 3, 4, 5})

	if s.Len() != 5 {
		t.Errorf("s.Len() = %d 而不是5", s.Len())
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
