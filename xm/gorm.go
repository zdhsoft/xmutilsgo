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

// 日期时间范围参数解析
//
//	 注意：开始时间和结束时间存在的时候：paramFieldName between ? and ?
//			都不存在的时候：不会生成 where 条件
//	     仅仅存在开始时间的时候：paramFieldName >= ?
//	     仅仅存在结束时间的时候：paramFieldName <= ?
//
//		SQL 语句的 where 条件的是日期时间字符串 YYYY-MM-DD HH:MM:SS
//		- paramFieldName: 时间字段名称
//		- paramBeginDateTime: 开始时间参数名称
//		- paramBeginTime: 开始时间字符串 (格式：YYYY-MM-DD HH:MM:SS)
//		- paramEndName: 结束时间参数名称
//		- paramEndDateTime: 结束时间字符串 (格式：YYYY-MM-DD HH:MM:SS)
//
// 返回值：
func (w *GormWhere) AddDateTimeScope(paramFieldName string, paramBeignName string, paramBeginDateTime string, paramEndName string, paramEndDateTime string) *BaseRet {
	r := NewBaseRet()
	for range [1]int{} {
		stBegin, retBegin := ParamDateTimeCheck(paramBeginDateTime, paramBeginDateTime, true)
		if retBegin.IsNotOK() {
			r.AssignErrorFrom(retBegin)
			break
		}
		stEnd, retEnd := ParamDateTimeCheck(paramEndDateTime, paramEndDateTime, true)
		if retEnd.IsNotOK() {
			r.AssignErrorFrom(retEnd)
			break
		}

		if stBegin != nil && stEnd != nil && stEnd.Before(*stBegin) {
			r.SetError(ERR_FAIL, paramEndName+"不能早于"+paramBeignName)
			break
		}
		// 除了 SQL 语句的 where 条件
		if stEnd == nil {
			if stBegin != nil {
				w.Add(paramFieldName+" >= ?", paramBeginDateTime)
			}
		} else {
			if stBegin != nil {
				w.Add(paramFieldName+" between ? and ?", paramBeginDateTime, paramEndDateTime)
			} else {
				w.Add(paramFieldName+" <= ?", paramEndDateTime)
			}
		}
	}
	return r
}

// 日期时间范围参数解析
//
//	 注意：开始时间和结束时间存在的时候：paramFieldName between ? and ?
//			都不存在的时候：不会生成 where 条件
//	     仅仅存在开始时间的时候：paramFieldName >= ?
//	     仅仅存在结束时间的时候：paramFieldName <= ?
//
//		SQL 语句的 where 条件的是日期时间字符串 YYYY-MM-DD
//		- paramFieldName: 时间字段名称
//		- paramBeignName: 开始时间参数名称
//		- paramBeginDate: 开始日期字符串 (格式：YYYY-MM-DD)
//		- paramEndName: 结束时间参数名称
//		- paramEndDate: 结束日期字符串 (格式：YYYY-MM-DD)
//
// 返回值：
func (w *GormWhere) AddDateScope(paramFieldName string, paramBeignName string, paramBeginDate string, paramEndName string, paramEndDate string) *BaseRet {
	r := NewBaseRet()
	for range [1]int{} {
		stBegin, retBegin := ParamDateCheck(paramBeginDate, paramBeginDate, true)
		if retBegin.IsNotOK() {
			r.AssignErrorFrom(retBegin)
			break
		}
		stEnd, retEnd := ParamDateCheck(paramEndDate, paramEndDate, true)
		if retEnd.IsNotOK() {
			r.AssignErrorFrom(retEnd)
			break
		}

		if stBegin != nil && stEnd != nil && stEnd.Before(*stBegin) {
			r.SetError(ERR_FAIL, paramEndName+"不能早于"+paramBeignName)
			break
		}
		// 除了 SQL 语句的 where 条件
		if stEnd == nil {
			if stBegin != nil {
				w.Add(paramFieldName+" >= ?", paramBeginDate)
			}
		} else {
			if stBegin != nil {
				w.Add(paramFieldName+" between ? and ?", paramBeginDate, paramEndDate)
			} else {
				w.Add(paramFieldName+" <= ?", paramEndDate)
			}
		}
	}
	return r
}

// 日期时间范围参数解析(时间戳)
//
//	 注意：开始时间和结束时间存在的时候：paramFieldName between ? and ?
//			都不存在的时候：不会生成 where 条件
//	     仅仅存在开始时间的时候：paramFieldName >= ?
//	     仅仅存在结束时间的时候：paramFieldName <= ?

//	SQL 语句的 where 条件的是时间戳 64位整型
//	- paramFieldName: 时间字段名称
//	- paramBeignName: 开始时间参数名称
//	- paramBeginDateTime: 开始时间字符串 (格式：YYYY-MM-DD HH:MM:SS)
//	- paramEndName: 结束时间参数名称
//	- paramEndDateTime: 结束时间字符串 (格式：YYYY-MM-DD HH:MM:SS)
//
// 返回值：
func (w *GormWhere) AddDateTimeScopeTimestamp(paramFieldName string, paramBeignName string, paramBeginDateTime string, paramEndName string, paramEndDateTime string) *BaseRet {
	r := NewBaseRet()
	for range [1]int{} {
		stBegin, retBegin := ParamDateTimeCheck(paramBeginDateTime, paramBeginDateTime, true)
		if retBegin.IsNotOK() {
			r.AssignErrorFrom(retBegin)
			break
		}
		stEnd, retEnd := ParamDateTimeCheck(paramEndDateTime, paramEndDateTime, true)
		if retEnd.IsNotOK() {
			r.AssignErrorFrom(retEnd)
			break
		}
		// 如果开始日期时间和结束日期时间都存在，则开始日期时间必须早于结束日期时间
		if stBegin != nil && stEnd != nil && stEnd.Before(*stBegin) {
			r.SetError(ERR_FAIL, paramEndName+"不能早于"+paramBeignName)
			break
		}
		// 除了 SQL 语句的 where 条件
		if stEnd == nil {
			if stBegin != nil {
				w.Add(paramFieldName+" >= ?", stBegin.Unix())
			}
		} else {
			if stBegin != nil {
				w.Add(paramFieldName+" between ? and ?", stBegin.Unix(), stEnd.Unix())
			} else {
				w.Add(paramFieldName+" <= ?", stEnd.Unix())
			}
		}
	}
	return r
}

// 日期范围参数解析()
//
//	  时间戳的单位是秒
//		 注意：开始时间和结束时间存在的时候：paramFieldName between ? and ?
//				都不存在的时候：不会生成 where 条件
//		     仅仅存在开始时间的时候：paramFieldName >= ?
//		     仅仅存在结束时间的时候：paramFieldName < ?
//		SQL 语句的 where 条件的是时间戳 64位整型
//		- paramFieldName: 时间字段名称
//		- paramBeignName: 开始时间参数名称
//		- paramBeginDate: 开始时间字符串 (格式：YYYY-MM-DD)
//		- paramEndName: 结束时间参数名称
//		- paramEndDate: 结束时间字符串 (格式：YYYY-MM-DD)
//
// 返回值：
func (w *GormWhere) AddDateScopeTimestamp(paramFieldName string, paramBeignName string, paramBeginDate string, paramEndName string, paramEndDate string) *BaseRet {
	r := NewBaseRet()
	for range [1]int{} {
		stBegin, retBegin := ParamDateCheck(paramBeginDate, paramEndDate, true)
		if retBegin.IsNotOK() {
			r.AssignErrorFrom(retBegin)
			break
		}
		stEnd, retEnd := ParamDateCheck(paramBeginDate, paramEndDate, true)
		if retEnd.IsNotOK() {
			r.AssignErrorFrom(retEnd)
			break
		}
		// 如果开始日期和结束日期都存在，则开始日期必须早于结束日期
		if stBegin != nil && stEnd != nil && stEnd.Before(*stBegin) {
			r.SetError(ERR_FAIL, paramEndName+"不能早于"+paramBeignName)
			break
		}
		// 除了 SQL 语句的 where 条件
		if stEnd == nil {
			if stBegin != nil {
				w.Add(paramFieldName+" >= ?", stBegin.Unix())
			}
		} else {
			if stBegin != nil {
				w.Add(paramFieldName+" between ? and ?", stBegin.Unix(), stEnd.Unix()+SECOND_BY_DAY-1)
			} else {
				w.Add(paramFieldName+" < ?", stEnd.Unix()+SECOND_BY_DAY)
			}
		}
	}
	return r
}

// 创建一个新的GormWhere对象
func NewGormWhere(cap int) *GormWhere {
	return &GormWhere{items: make([]*whereItem, 0, cap)}
}
