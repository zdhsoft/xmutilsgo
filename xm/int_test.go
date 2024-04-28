package xm

import (
	"testing"
)

func Test_I(t *testing.T) {
	if I(10) != "10" {
		t.Errorf("I(%d) != 10 failed", 10)
	}
	if I(20) != "20" {
		t.Errorf("I(%d) != 20 failed", 20)
	}
	if I(-1) != "-1" {
		t.Errorf("I(%d) != -1 failed", -1)
	}

	if I(0x7fffffff) != "2147483647" {
		t.Errorf("I(%d) != 2147483647 failed", 0x7fffffff)
	}
	if I(-0x80000000) != "-2147483648" {
		t.Errorf("I(%d) != -2147483648 failed", -0x80000000)
	}

	if I(0x100000000) != "4294967296" {
		t.Errorf("I(%d) != 4294967296 failed", 0x100000000)
	}
	if I(-0x100000000) != "-4294967296" {
		t.Errorf("I(%d) != -4294967296 failed", -0x100000000)
	}

	if I(0x7fffffffffffffff) != "9223372036854775807" {
		t.Errorf("I(%d) != 9223372036854775807 failed", 0x7fffffffffffffff)
	}
	if I(-0x7fffffffffffffff) != "-9223372036854775807" {
		t.Errorf("I(%d) != -9223372036854775807 failed", -0x7fffffffffffffff)
	}
}
func Test_Int2String(t *testing.T) {
	if Int2String(10) != "10" {
		t.Errorf("Int2String(%d) != 10 failed", 10)
	}
	if Int2String(20) != "20" {
		t.Errorf("Int2String(%d) != 20 failed", 20)
	}
	if Int2String(-1) != "-1" {
		t.Errorf("Int2String(%d) != -1 failed", -1)
	}

	if Int2String(0x7fffffff) != "2147483647" {
		t.Errorf("Int2String(%d) != 2147483647 failed", 0x7fffffff)
	}
	if Int2String(-0x80000000) != "-2147483648" {
		t.Errorf("Int2String(%d) != -2147483648 failed", -0x80000000)
	}

	if Int2String(0x100000000) != "4294967296" {
		t.Errorf("Int2String(%d) != 4294967296 failed", 0x100000000)
	}
	if Int2String(-0x100000000) != "-4294967296" {
		t.Errorf("Int2String(%d) != -4294967296 failed", -0x100000000)
	}

	if Int2String(0x7fffffffffffffff) != "9223372036854775807" {
		t.Errorf("Int2String(%d) != 9223372036854775807 failed", 0x7fffffffffffffff)
	}
	if Int2String(-0x7fffffffffffffff) != "-9223372036854775807" {
		t.Errorf("Int2String(%d) != -9223372036854775807 failed", -0x7fffffffffffffff)
	}
}

func Test_U(t *testing.T) {
	if U(uint32(10)) != "10" {
		t.Errorf("U(%d) != 10 failed", 10)
	}
	if U(uint64(20)) != "20" {
		t.Errorf("U(%d) != 20 failed", 20)
	}
	if U(uint32(0xffffffff)) != "4294967295" {
		t.Errorf("U(%d) != -1 failed", -1)
	}

	if U(uint32(0x7fffffff)) != "2147483647" {
		t.Errorf("U(%d) != 2147483647 failed", uint32(0x7fffffff))
	}
	if U(uint32(0x80000000)) != "2147483648" {
		t.Errorf("U(%d) != 2147483648 failed", uint32(0x80000000))
	}

	if U(uint64(0x100000000)) != "4294967296" {
		t.Errorf("U(%d) != 4294967296 failed", uint64(0x100000000))
	}
	if U(uint64(0x100000000)) != "4294967296" {
		t.Errorf("U(%d) != 4294967296 failed", uint64(0x100000000))
	}

	if U(uint64(0x7fffffffffffffff)) != "9223372036854775807" {
		t.Errorf("U(%d) != 9223372036854775807 failed", uint64(0x7fffffffffffffff))
	}
	u := uint64(0xffffffffffffffff)
	if U(u) != "18446744073709551615" {
		t.Errorf("U(%d) != 18446744073709551615 failed ==>%s", uint64(u), U(u))
	}
}
