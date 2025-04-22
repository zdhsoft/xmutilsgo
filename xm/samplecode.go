package xm

import "fmt"

const (
	randomMultiplier = 0x015a4e35
	randomIncrement  = 1
)

// 编解码结构
type SimpleCode struct {
	seed int
	fix  int
}

// 检查参数合法性
func checkParam(paramData []byte, paramMinLen int) error {
	if paramData == nil {
		return fmt.Errorf("data is nil")
	}
	if len(paramData) < paramMinLen {
		return fmt.Errorf("data size is too low")
	}
	return nil
}

// 伪随机数生成器
func (s *SimpleCode) rand() int {
	m := randomMultiplier*s.seed + randomIncrement + s.fix
	s.seed = m & 0x7fffffff
	s.fix = (s.fix + 1) % 0x7fffffff
	return s.seed
}

// 初始化种子
func (s *SimpleCode) SetSeed(paramSeed int) {
	s.seed = (paramSeed & 0x7fffffff) % 256
	s.fix = s.seed
}

// 获取种子
func (s *SimpleCode) GetSeed() int {
	return s.seed
}

// Encode 加密函数
func (s *SimpleCode) Encode(paramData []byte) ([]byte, error) {
	err := checkParam(paramData, 0)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(paramData))
	for i := 0; i < len(paramData); i++ {
		value := s.rand()%256 + int(paramData[i])
		result[i] = byte(value % 256)
	}
	return result, nil
}

// Decode 解密函数
func (s *SimpleCode) Decode(paramData []byte) ([]byte, error) {
	err := checkParam(paramData, 0)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(paramData))
	for i := 0; i < len(paramData); i++ {
		value := s.rand() % 256
		v := int(paramData[i]) - value
		if v < 0 {
			v = (v + 256) % 256
		}
		result[i] = byte(v)
	}
	return result, nil
}

// EncodePackage 将 seed 放在第一字节
func (s *SimpleCode) EncodePackage(paramData []byte) ([]byte, error) {
	err := checkParam(paramData, 0)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(paramData)+1)
	result[0] = byte(s.seed) // 保存初始种子
	for i := 0; i < len(paramData); i++ {
		value := s.rand()%256 + int(paramData[i])
		result[i+1] = byte(value % 256)
	}
	return result, nil
}

// DecodePackage 解析 EncodePackage 生成的包
func (s *SimpleCode) DecodePackage(paramData []byte) ([]byte, error) {
	err := checkParam(paramData, 0)
	if err != nil {
		return nil, err
	}

	seed := int(paramData[0])
	s.SetSeed(seed)

	result := make([]byte, len(paramData)-1)
	for i := 0; i < len(result); i++ {
		value := s.rand() % 256
		v := int(paramData[i+1]) - value
		if v < 0 {
			v = (v + 256) % 256
		}
		result[i] = byte(v)
	}
	return result, nil
}

// NewSampleCode 创建一个samplecode实例
func NewSampleCode(paramSeed int) *SimpleCode {
	s := &SimpleCode{}
	s.SetSeed(paramSeed)
	return s
}

// NewSampleCodeByMillis 创建一个samplecode实例，种子为当前时间戳
func NewSampleCodeByMillis() *SimpleCode {
	return NewSampleCode(int(GetNowMillis()))
}
