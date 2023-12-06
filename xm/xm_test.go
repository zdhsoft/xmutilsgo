package xm

import (
	"testing"
)

func TestInt(t *testing.T) {

	v :=Int2StringPad(1000, 20)
	if v != "00000000000000001000" {
		t.Errorf("Int2StringPad(1000, 20) != 00000000000000001000")
	}
}
