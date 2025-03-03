package xm

import "time"

// 参数处理检验的公共函数

// 日期时间参数解析
//   - paramName: 参数名称
//   - paramDateTime: 日期时间字符串 (格式：YYYY-MM-DD hh:mm:ss)
//   - paramAllocEmpty: 是否允许空字符串 (如果paramDateTime是空的字符串，则返回 nil)
//   - paramR: 错误信息对象
//   - 返回值：*time.Time 日期时间对象
//   - 如果解析失败，返回 nil
//   - 如果解析成功，返回 *time.Time 对象
func ParamDateTimeCheck(paramName string, paramDateTime string, paramAllocEmpty bool) (*time.Time, *BaseRet) {
	r := NewBaseRet()
	var resultRet *time.Time = nil
	for range [1]int{} {
		if paramDateTime == "" {
			if !paramAllocEmpty {
				r.SetError(ERR_FAIL, paramName+"参数不能为空")
			}
			break
		}
		if !IsDateTimeFormat(paramDateTime) {
			r.SetError(ERR_FAIL, paramName+"参数格式错误，应为YYYY-MM-DD hh:mm:ss")
			break
		}
		result, err := ParseDateTimeForBeijingTime(paramDateTime)
		if err != nil {
			r.SetError(ERR_FAIL, paramName+" = "+paramDateTime+" 是一个无效日期时间:"+err.Error())
			break
		}
		resultRet = &result
	}
	return resultRet, r
}

// 日期参数解析
//   - paramName: 参数名称
//   - paramDate: 日期字符串 (格式：YYYY-MM-DD)
//   - paramAllocEmpty: 是否允许空字符串 (如果paramDate是空的字符串，则返回 nil)
//   - paramR: 错误信息对象
//   - 返回值：*time.Time 日期时间对象
//   - 如果解析失败，返回 nil
//   - 如果解析成功，返回 *time.Time 对象
func ParamDateCheck(paramName string, paramDate string, paramAllocEmpty bool) (*time.Time, *BaseRet) {
	r := NewBaseRet()
	var resultRet *time.Time = nil
	for range [1]int{} {
		if paramDate == "" {
			if !paramAllocEmpty {
				r.SetError(ERR_FAIL, paramName+"参数不能为空")
			}
			break
		}
		if !IsDateFormat(paramDate) {
			r.SetError(ERR_FAIL, paramName+"参数格式错误，应为YYYY-MM-DD")
			break
		}
		result, err := ParseDateForBeijingTime(paramDate)
		if err != nil {
			r.SetError(ERR_FAIL, paramName+" = "+paramDate+" 是一个无效日期:"+err.Error())
			break
		}
		resultRet = &result
	}
	return resultRet, r
}

// 日期或日期时间参数解析
//   - paramName: 参数名称
//   - paramDateTimeOrDate: 日期时间字符串 (格式：YYYY-MM-DD或YYYY-MM-DD hh:mm:ss)
//   - paramAllocEmpty: 是否允许空字符串 (如果paramDateTimeOrDate是空的字符串，则返回 nil)
//   - paramR: 错误信息对象
//   - 返回值：*time.Time 日期时间对象
//   - 如果解析失败，返回 nil
//   - 如果解析成功，返回 *time.Time 对象
func ParamDateOrDateTimeCheck(paramName string, paramDateTimeOrDate string, paramAllocEmpty bool) (*time.Time, *BaseRet) {
	r := NewBaseRet()
	var resultRet *time.Time = nil
	for range [1]int{} {
		if paramDateTimeOrDate == "" {
			if !paramAllocEmpty {
				r.SetError(ERR_FAIL, paramName+"参数不能为空")
			}
			break
		}
		if !(IsDateFormat(paramDateTimeOrDate) || IsDateTimeFormat(paramDateTimeOrDate)) {
			r.SetError(ERR_FAIL, paramName+"参数格式错误，应为YYYY-MM-DD或YYYY-MM-DD hh:mm:ss")
			break
		}
		result, err := ParseDateTimeForBeijingTime(paramDateTimeOrDate)
		if err != nil {
			r.SetError(ERR_FAIL, paramName+" = "+paramDateTimeOrDate+" 是一个无效日期或日期时间:"+err.Error())
			break
		}
		resultRet = &result
	}
	return resultRet, r
}
