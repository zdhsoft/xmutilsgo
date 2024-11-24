package xm

import (
	"fmt"
	"math"
)

// GetMapKeys 取Map的Key数组
func GetMapKeys[K comparable, V any](paramMap map[K]V) []K {
	keys := make([]K, 0, len(paramMap))
	for key := range paramMap {
		keys = append(keys, key)
	}
	return keys
}

// GetMapValues 取Map的Value数组
func GetMapValues[K comparable, V any](paramMap map[K]V) []V {
	Values := make([]V, 0, len(paramMap))
	for _, v := range paramMap {
		Values = append(Values, v)
	}
	return Values
}

// GetMapKeyValue 取Map的KeyValue数组
func GetMapKeyValue[K comparable, V any](paramMap map[K]V) ([]K, []V) {
	Values := make([]V, 0, len(paramMap))
	keys := make([]K, 0, len(paramMap))
	for k, v := range paramMap {
		Values = append(Values, v)
		keys = append(keys, k)
	}
	return keys, Values
}

// 计算百分比
//   - num: 已完成的数量
//   - total: 总数量
//   - 返回值是百分比字符串，保留两位小数(返回值不含%)
func CalcPercent(num, total int64) string {
	if total == 0 || num == 0 {
		return "0"
	}
	percent := float64(num) / float64(total) * 10000
	rndPercent := math.Round(percent) / 100
	return fmt.Sprintf("%.2f", rndPercent)
}

// 将单位为分转换为浮点的元
//   - 单位是分
func CalcMoneyByFloat(amount float64) string {
	return CalcMoneyByCent(int64(math.Round(amount * 100)))
}

// 将单位为分转换为浮点的元，保留两位小数, 不足补0
//   - 单位是分
func CalcMoneyByFloat2(amount float64) string {
	return CalcMoneyByCent2(int64(math.Round(amount * 100)))
}

// 将单位为分转换为浮点的元
//   - 单位是分
func CalcMoneyByCent(num int64) string {
	// return fmt.Sprintf("%.2f", float64(num)/100)
	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}
	fen := num % 100
	cnt := num / 100
	if !isNegative {
		if fen > 0 {
			if fen%10 == 0 {
				return fmt.Sprintf("%d.%d", cnt, fen/10)
			} else {
				return fmt.Sprintf("%d.%02d", cnt, fen)
			}
		} else {
			// 如果没有小数位，则显示整数
			return fmt.Sprintf("%d", cnt)
		}
	} else {
		if fen > 0 {
			if fen%10 == 0 {
				return fmt.Sprintf("-%d.%d", cnt, fen/10)
			} else {
				return fmt.Sprintf("-%d.%02d", cnt, fen)
			}
		} else {
			// 如果没有小数位，则显示整数
			return fmt.Sprintf("-%d", cnt)
		}
	}
}

// 将单位为分转换为浮点的元, 保留两位小数, 不足补0
//   - 单位是分
func CalcMoneyByCent2(num int64) string {
	// return fmt.Sprintf("%.2f", float64(num)/100)
	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}
	fen := num % 100
	cnt := num / 100
	if !isNegative {
		if fen > 0 {
			if fen%10 == 0 {
				return fmt.Sprintf("%d.%d0", cnt, fen/10)
			} else {
				return fmt.Sprintf("%d.%02d", cnt, fen)
			}
		} else {
			// 如果没有小数位，则显示整数
			return fmt.Sprintf("%d.00", cnt)
		}
	} else {
		if fen > 0 {
			if fen%10 == 0 {
				return fmt.Sprintf("-%d.%d0", cnt, fen/10)
			} else {
				return fmt.Sprintf("-%d.%02d", cnt, fen)
			}
		} else {
			// 如果没有小数位，则显示整数
			return fmt.Sprintf("-%d.00", cnt)
		}
	}
}
