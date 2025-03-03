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

func Test_WhereTimeScope(t *testing.T) {

	strDate1 := "2021-01-01"
	strDate2 := "2021-01-31"

	w1 := GormWhere{}
	r1 := w1.AddDateScope("test_date", "begin_date", strDate1, "end_date", strDate2)
	if r1.IsNotOK() {
		t.Errorf("AddDateScope() = %s, ret=%d, want 0", r1.Msg, r1.Ret)
	}
	if w1.WhereString() != "test_date between ? and ?" {
		t.Errorf("WhereString() = %s, want %s", w1.WhereString(), "test_date between ? and ?")
	}

	values1 := w1.WhereValues()
	if len(values1) != 2 {
		t.Errorf("len(Values()) = %d, want %v", len(values1), 2)
	}

	if values1[0] != strDate1 {
		t.Errorf("Values(0) = %s, want %d", strDate1, values1[0])
	}

	if values1[1] != strDate2 {
		t.Errorf("Values(1) = %s, want %d", strDate2, values1[1])
	}
}
