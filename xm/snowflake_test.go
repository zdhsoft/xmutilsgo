package xm

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {

	// 创建雪花算法实例，workerID为1
	sf, err := NewSnowflake(1)
	if err != nil {
		t.Fatal(err)
	}

	// 生成ID
	id := sf.NextID()
	fmt.Printf("生成的ID: %d\n", id)

	// 获取分表键（假设有16个分表）
	shardKey := GetShardKey(id, 16)
	fmt.Printf("分表键: %d\n", shardKey)

	// 生成多个ID并查看分表分布
	shardCount := 16
	shardDistribution := make(map[int]int)
	for i := 0; i < 1000; i++ {
		id := sf.NextID()
		fmt.Printf("生成的ID: %d\n", id)
		shardKey := GetShardKey(id, shardCount)
		shardDistribution[shardKey]++
	}

	// 打印分表分布情况
	fmt.Println("\n分表分布情况:")
	for i := 0; i < shardCount; i++ {
		fmt.Printf("分表 %d: %d 条数据\n", i, shardDistribution[i])
	}
}
