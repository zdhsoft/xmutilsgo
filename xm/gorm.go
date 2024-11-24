package xm

//  这个提供的工具，用于简化gorm的where条件的生成。
import "strings"

type whereItem struct {
	Cond  string        // 条件字符串
	Value []interface{} // 条件值列表
}

// Gorm的Where类
type GormWhere struct {
	items []*whereItem
}

// 添加条件
func (w *GormWhere) Add(cond string, value ...interface{}) {
	w.items = append(w.items, &whereItem{Cond: cond, Value: value})
}

// 条件数量
func (w *GormWhere) Len() int {
	return len(w.items)
}

// 生成where条件语句和参数列表
func (w *GormWhere) Where() (string, []interface{}) {
	if len(w.items) == 0 {
		return "", []interface{}{}
	}
	builder := &strings.Builder{}
	builder.Grow(1024)

	values := make([]interface{}, 0, len(w.items)*2)
	first := true
	for _, item := range w.items {
		if first {
			first = false
		} else {
			builder.WriteString(" and ")
		}
		builder.WriteString(item.Cond)
		values = append(values, item.Value...)
	}
	return builder.String(), values
}

// 生成where条件语句
func (w *GormWhere) WhereString() string {
	if len(w.items) == 0 {
		return ""
	}
	builder := &strings.Builder{}
	builder.Grow(1024)
	first := true
	for _, item := range w.items {
		if first {
			first = false
		} else {
			builder.WriteString(" and ")
		}
		builder.WriteString(item.Cond)
	}
	return builder.String()
}

// 生成where条件参数列表
func (w *GormWhere) WhereValues() []interface{} {
	if len(w.items) == 0 {
		return []interface{}{}
	}
	values := make([]interface{}, 0, len(w.items)*2)
	for _, item := range w.items {
		values = append(values, item.Value...)
	}
	return values
}

// 创建一个新的GormWhere对象
func NewGormWhere(cap int) *GormWhere {
	return &GormWhere{items: make([]*whereItem, 0, cap)}
}
