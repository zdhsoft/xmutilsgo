package xm

import (
	"errors"
	"sync"
	"time"
)

// 雪花算法实现：
//   使用64位整数存储ID
//   包含时间戳（41位）、机器ID（10位）和序列号（12位）
//   支持分布式环境下的唯一性
//   时间戳从2025年开始，可以使用约69年
// 分表策略：
//   使用 GetShardKey 函数根据ID计算分表键
//   采用取模运算，确保数据均匀分布
//   分表数量建议为2的幂次方（如16、32、64等）

const (
	// 雪花算法开始时间戳 (2025-01-01 00:00:00)
	Epoch = 1735660800000
	// 机器ID所占的位数
	workerIDBits = 10
	// 序列号所占的位数
	sequenceBits = 12
	// 支持的最大机器ID 1023
	MaxWorkerID = -1 ^ (-1 << workerIDBits)
	// 支持的最大序列号 4095
	maxSequence = -1 ^ (-1 << sequenceBits)
	// 机器ID向左移位数
	workerIDShift = sequenceBits
	// 时间戳向左移位数
	timestampShift = sequenceBits + workerIDBits
)

// Snowflake 雪花算法结构体
type Snowflake struct {
	mu        sync.Mutex
	workerID  int64
	sequence  int64
	lastStamp int64
}

// NewSnowflake 创建一个新的雪花算法实例
func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > MaxWorkerID {
		return nil, ErrInvalidWorkerID
	}
	return &Snowflake{
		workerID: workerID,
	}, nil
}

// NextID 生成下一个ID
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 当前时间戳（毫秒）

	// 如果当前时间小于上次生成ID的时间，说明发生了时钟回拨
	if now < s.lastStamp {
		// 等待时钟追上
		time.Sleep(time.Duration(s.lastStamp-now) * time.Millisecond)
		now = time.Now().UnixNano() / 1e6
	}

	// 如果是同一毫秒内生成的
	if now == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSequence
		// 如果序列号用完，等待下一毫秒
		if s.sequence == 0 {
			// 使用time.Sleep替代空转
			time.Sleep(time.Millisecond)
			now = time.Now().UnixNano() / 1e6
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = now

	// 生成ID
	id := ((now - Epoch) << timestampShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}

// GetShardKey 根据ID获取分表键
func GetShardKey(id int64, shardCount int) int {
	// 使用ID的低位进行分表，确保数据分布均匀
	return int(id & int64(shardCount-1))
}

// 错误定义
var (
	ErrInvalidWorkerID = errors.New("invalid worker ID")
)
