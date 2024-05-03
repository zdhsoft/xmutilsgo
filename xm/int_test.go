package xm

import (
	"testing"
)

func Test_I(t *testing.T) {
	v := []int64{10, 20, -1, 0x7fffffff, -0x80000000, 0x100000000, 0x7fffffffffffffff, -0x7fffffffffffffff, 2147483647, -2147483648, 4294967296, 9223372036854775807, -9223372036854775807}
	dest := []string{"10", "20", "-1", "2147483647", "-2147483648", "4294967296", "9223372036854775807", "-9223372036854775807", "2147483647", "-2147483648", "4294967296", "9223372036854775807", "-9223372036854775807"}
	for i, val := range v {
		if I(val) != dest[i] {
			t.Errorf("I(%d) != %s failed", val, dest[i])
		}
	}
}

func Test_Int2String(t *testing.T) {
	v := []int64{10, 20, -1, 0x7fffffff, -0x80000000, 0x100000000, 0x7fffffffffffffff, -0x7fffffffffffffff, 2147483647, -2147483648, 4294967296, 9223372036854775807, -9223372036854775807}
	dest := []string{"10", "20", "-1", "2147483647", "-2147483648", "4294967296", "9223372036854775807", "-9223372036854775807", "2147483647", "-2147483648", "4294967296", "9223372036854775807", "-9223372036854775807"}
	for i, val := range v {
		if Int2String(val) != dest[i] {
			t.Errorf("I(%d) != %s failed", val, dest[i])
		}
	}
}

func Test_U(t *testing.T) {
	v := []uint64{10, 20, 0xffffffff, 0x7fffffff, 0x80000000, 0x100000000, 0x7fffffffffffffff, 0xffffffffffffffff, 0x7fffffff, 0x80000000, 0x100000000, 0x7fffffffffffffff, 0xffffffffffffffff, 2147483647, 2147483648, 4294967296, 9223372036854775807, 18446744073709551615}
	dest := []string{"10", "20", "4294967295", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615"}
	for i, val := range v {
		if U(val) != dest[i] {
			t.Errorf("U(%d) != %s failed", val, dest[i])
		}
	}
}

func Test_Uint2String(t *testing.T) {
	v := []uint64{10, 20, 0xffffffff, 0x7fffffff, 0x80000000, 0x100000000, 0x7fffffffffffffff, 0xffffffffffffffff, 0x7fffffff, 0x80000000, 0x100000000, 0x7fffffffffffffff, 0xffffffffffffffff, 2147483647, 2147483648, 4294967296, 9223372036854775807, 18446744073709551615}
	dest := []string{"10", "20", "4294967295", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615", "2147483647", "2147483648", "4294967296", "9223372036854775807", "18446744073709551615"}
	for i, val := range v {
		if UInt2String(val) != dest[i] {
			t.Errorf("U(%d) != %s failed", val, dest[i])
		}
	}
}

func Test_RandomIntScope(t *testing.T) {
	minValue := int(100)
	maxValue := int(10000)

	for i := 0; i < 1000; i++ {
		val := RandomIntScope(minValue, maxValue)
		if val < minValue || val > maxValue {
			t.Errorf("RandomIntScope(%d, %d) failed", minValue, maxValue)
		}
	}
}

func Test_ReverseInt64(t *testing.T) {
	v := []int64{123456789, 987654321, 123456789012345678, 9223372036854775807, -9223372036854775807}
	dest := []int64{987654321, 123456789, 876543210987654321, 7085774586302733229, -7085774586302733229}
	for i, val := range v {
		if ReverseInt64(val) != dest[i] {
			t.Errorf("ReverseInt64(%d) => %d != %d failed", val, ReverseInt64(val), dest[i])
		}
	}
}

func Test_RandOne(t *testing.T) {
	v := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 1000; i++ {
		val, err := RandOneInArray(v)
		if err != nil {
			t.Errorf("RandOneInArray(%v) failed: %s", v, err.Error())
		}
		if *val < 1 || *val > 10 {
			t.Errorf("RandOneInArray(%v)  value=%d  failed", v, *val)
		}
	}
}
