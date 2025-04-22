package xm

import (
	"testing"
)

func TestSampleCode(t *testing.T) {
	sc := NewSampleCodeByMillis()
	// 示例数据
	data := []byte("Hello SimpleCode")
	seed := sc.GetSeed()
	// fmt.Println("Seed:", seed)
	// Encode
	encoded, err := sc.Encode(data)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	// fmt.Println("Encoded:", encoded)

	// Decode
	sc.SetSeed(seed)
	decoded, err := sc.Decode(encoded)
	if err != nil {
		t.Errorf("Decode error: %v", err)
	}
	// fmt.Println("Decoded:", string(decoded))

	if string(decoded) != string(data) {
		t.Errorf("Decode error: %v", err)
	}

	// EncodePackage
	sc.SetSeed(91234)
	pack, err := sc.EncodePackage(data)
	if err != nil {
		t.Errorf("EncodePackage error: %v", err)
	}
	// fmt.Println("Package:", pack)

	// DecodePackage
	unpack, err := sc.DecodePackage(pack)
	if err != nil {
		t.Errorf("DecodePackage error: %v", err)
	}
	// fmt.Println("Unpacked:", string(unpack))
	if string(unpack) != string(data) {
		t.Errorf("DecodePackage error: %v", err)
	}
}
