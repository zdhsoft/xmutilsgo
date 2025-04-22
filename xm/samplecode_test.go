package xm

import (
	"fmt"
	"testing"
)

func TestSampleCode(t *testing.T) {
	sc := &SimpleCode{}
	// 示例数据
	data := []byte("Hello SimpleCode")
	seed := int(GetNowMillis())
	fmt.Println("Seed:", seed)
	// Encode
	encoded := sc.Encode(seed, data)
	fmt.Println("Encoded:", encoded.Data)

	// Decode
	decoded := sc.Decode(seed, encoded.Data)
	fmt.Println("Decoded:", string(decoded.Data))

	// EncodePackage
	pack := sc.EncodePackage(seed, data)
	fmt.Println("Package:", pack.Data)

	// DecodePackage
	unpack := sc.DecodePackage(pack.Data)
	fmt.Println("Unpacked:", string(unpack.Data))
}
