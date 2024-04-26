package xm

// 取Map的Key数组
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// 取Map的Value数组
func GetMapValues[K comparable, V any](m map[K]V) []V {
	Values := make([]V, 0, len(m))
	for _, v := range m {
		Values = append(Values, v)
	}
	return Values
}

// 取Map的KeyValue数组
func GetMapKeyValue[K comparable, V any](m map[K]V) ([]K, []V) {
	Values := make([]V, 0, len(m))
	keys := make([]K, 0, len(m))
	for k, v := range m {
		Values = append(Values, v)
		keys = append(keys, k)
	}
	return keys, Values
}
