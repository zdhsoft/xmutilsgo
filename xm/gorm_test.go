package xm

import "testing"

func Test_Where(t *testing.T) {
	stWhere := GormWhere{}
	stWhere.Add("id = ?", 100)
	stWhere.Add("name = ?", "test")
	stWhere.Add("between ? and ?", 10, 20)

	dest_value := []interface{}{100, "test", 10, 20}
	dest_str := "id = ? and name = ? and between ? and ?"
	if stWhere.WhereString() != dest_str {
		t.Errorf("WhereString() = %s, want %s", stWhere.WhereString(), dest_str)
	}
	if IsEqualArray[interface{}](stWhere.WhereValues(), dest_value) == false {
		t.Errorf("Values() = %v, want %v", stWhere.WhereValues(), dest_value)
	}

	w, v := stWhere.Where()
	if w != dest_str {
		t.Errorf("Where() = %s, want %s", w, dest_str)
	}
	if IsEqualArray[interface{}](v, dest_value) == false {
		t.Errorf("Values() = %v, want %v", v, dest_value)
	}

	if stWhere.Len() != 3 {
		t.Errorf("Len() = %d, want %d", stWhere.Len(), 3)
	}
}
