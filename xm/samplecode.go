package xm

const (
	randomMultiplier = 0x015a4e35
	randomIncrement  = 1
)

// 错误码定义
type SimpleCodeError int

const (
	CodeOK             SimpleCodeError = 0
	CodeSeedNotInt     SimpleCodeError = 1
	CodeDataNotBuffer  SimpleCodeError = 2
	CodeDataSizeTooLow SimpleCodeError = 3
)

// 返回结构
type SimpleCodeResult struct {
	Code SimpleCodeError
	Seed int
	Data []byte
}

// 编解码结构
type SimpleCode struct {
	seed int
	fix  int
}

// 检查参数合法性
func checkParam(seed int, data []byte, minLen int) SimpleCodeError {
	if data == nil {
		return CodeDataNotBuffer
	}
	if len(data) < minLen {
		return CodeDataSizeTooLow
	}
	return CodeOK
}

// 伪随机数生成器（非密码学安全）
func (s *SimpleCode) rand() int {
	m := randomMultiplier*s.seed + randomIncrement + s.fix
	s.seed = m & 0x7fffffff
	s.fix = (s.fix + 1) % 0x7fffffff
	return s.seed
}

// 初始化种子
func (s *SimpleCode) setSeed(seed int) {
	s.seed = (seed & 0x7fffffff) % 256
	s.fix = s.seed
}

// Encode 加密函数
func (s *SimpleCode) Encode(seed int, data []byte) SimpleCodeResult {
	if r := checkParam(seed, data, 0); r != CodeOK {
		return SimpleCodeResult{Code: r}
	}
	s.setSeed(seed)

	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		value := s.rand()%256 + int(data[i])
		result[i] = byte(value % 256)
	}
	return SimpleCodeResult{Code: CodeOK, Seed: seed, Data: result}
}

// Decode 解密函数
func (s *SimpleCode) Decode(seed int, data []byte) SimpleCodeResult {
	if r := checkParam(seed, data, 0); r != CodeOK {
		return SimpleCodeResult{Code: r}
	}
	s.setSeed(seed)

	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		value := s.rand() % 256
		v := int(data[i]) - value
		if v < 0 {
			v = (v + 256) % 256
		}
		result[i] = byte(v)
	}
	return SimpleCodeResult{Code: CodeOK, Seed: seed, Data: result}
}

// EncodePackage 将 seed 放在第一字节
func (s *SimpleCode) EncodePackage(seed int, data []byte) SimpleCodeResult {
	if r := checkParam(seed, data, 0); r != CodeOK {
		return SimpleCodeResult{Code: r}
	}
	s.setSeed(seed)

	result := make([]byte, len(data)+1)
	result[0] = byte(s.seed) // 保存初始种子
	for i := 0; i < len(data); i++ {
		value := s.rand()%256 + int(data[i])
		result[i+1] = byte(value % 256)
	}
	return SimpleCodeResult{Code: CodeOK, Seed: seed, Data: result}
}

// DecodePackage 解析 EncodePackage 生成的包
func (s *SimpleCode) DecodePackage(data []byte) SimpleCodeResult {
	if r := checkParam(0, data, 1); r != CodeOK {
		return SimpleCodeResult{Code: r}
	}
	seed := int(data[0])
	s.setSeed(seed)

	result := make([]byte, len(data)-1)
	for i := 0; i < len(result); i++ {
		value := s.rand() % 256
		v := int(data[i+1]) - value
		if v < 0 {
			v = (v + 256) % 256
		}
		result[i] = byte(v)
	}
	return SimpleCodeResult{Code: CodeOK, Seed: seed, Data: result}
}
