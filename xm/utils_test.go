package xm

import (
	"testing"
)

func Test_CalcMoneyByCent2(t *testing.T) {
	v := []int64{123, 456, 789, 120, 130, 131, 140, 151, 109}
	dest := []string{"1.23", "4.56", "7.89", "1.20", "1.30", "1.31", "1.40", "1.51", "1.09"}
	dest1 := []string{"1.23", "4.56", "7.89", "1.2", "1.3", "1.31", "1.4", "1.51", "1.09"}
	for i, val := range v {
		result2 := CalcMoneyByCent2(val)
		result := CalcMoneyByCent(val)
		if result2 != dest[i] {
			t.Errorf("CalcMoneyByCent2(%d) = %s, want %s", val, result2, dest[i])
		} else {
			// t.Logf("CalcMoneyByCent2(%d) => %s", val, result2)
		}
		if result != dest1[i] {
			t.Errorf("CalcMoneyByCent(%d) = %s, want %s", val, result, dest1[i])
		}
	}
}

// // 将浮点数分解为 >=0.01 和 <0.01 的两部分
// func splitDecimal(value float64) (greaterEqual int64, lessThan float64) {
// 	// 使用 decimal 库转换输入值
// 	d := decimal.NewFromFloat(value)
// 	d = d.Mul(decimal.NewFromFloat(100))
// 	// 将值截断到两位小数，得到 >= 0.01 的部分
// 	greaterEqual = d.IntPart()
// 	// lessThan 是 value 中小于 0.01 的部分
// 	lessThan, _ = d.Sub(d.Truncate(0)).Float64()
// 	return
// }

// func Test_V(t *testing.T) {
// 	// 测试不同的值
// 	values := []float64{2.3, 2.345, 3.14159, 5.678}
// 	for _, v := range values {
// 		greaterEqual, lessThan := splitDecimal(v)
// 		t.Logf("Original Value: %.5f\n", v)
// 		t.Logf(">= 0.01 Part: %d\n", greaterEqual)
// 		t.Logf("< 0.01 Part: %f\n\n", lessThan)
// 	}
// }

func Test_StringSpace(t *testing.T) {
	s := "hello world"
	s1 := "hello world\n k l adf"
	if !HasWhitespace(s) {
		t.Errorf("HasWhitespace(%s) = false, want true", s)
	}
	if RemoveAllWhiteSpace(s) != "helloworld" {
		t.Errorf("RemoveAllWhiteSpace(%s) = %s, want helloworld", s, RemoveAllWhiteSpace(s))
	}
	if !HasWhitespace(s1) {
		t.Errorf("HasWhitespace(%s) = false, want true", s1)
	}
	if RemoveAllWhiteSpace(s1) != "helloworldkladf" {
		t.Errorf("RemoveAllWhiteSpace(%s) = %s, want helloworldkladf", s1, RemoveAllWhiteSpace(s1))
	}
}
