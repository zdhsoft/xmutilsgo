package xm

import "fmt"

type PageInfo struct {
	// 第几页，从1开始
	Page int `json:"page"`
	// 每页的记录数，等同于length
	PageSize int `json:"page_size"`
}

// NewPage 创建新的Page对象
func NewPage[T Integer](paramPage T, paramPageSize T) *PageInfo {
	p := PageInfo{
		Page:     int(paramPage),
		PageSize: int(paramPageSize),
	}
	return p.RoundPageInfo()
}

// NewPageRoundSize 创建新的Page对象, 并围绕页大小 使其在合理值的范围内(指定缺省页值和最大值)
//   - paramPage 页码 (从1开始)
//   - paramPageSize 页大小 (每页记录数)
//   - paramDefaultSize 缺省页大小 (当paramPageSize <= 0时, 页大小取缺省值)
//   - paramMaxSize 最大页大小 (当paramPageSize > paramMaxSize时, 页大小取最大值)
func NewPageRoundSize[T Integer](paramPage T, paramPageSize T, paramDefaultSize T, paramMaxSize T) *PageInfo {
	p := PageInfo{
		Page:     int(paramPage),
		PageSize: int(paramPageSize),
	}
	return p.RoundPageInfoEx(int(paramDefaultSize), int(paramMaxSize))
}

// RoundPageInfo 围绕页信息 使它在合理值的范围内
func (p *PageInfo) RoundPageInfo() *PageInfo {
	return p.RoundPageInfoEx(10, 1000)
}

/*
RoundPageInfoEx 围绕页信息 使它在合理值的范围内(指定缺省页值和最大值)
  - paramDefaultSize 缺省页大小
  - paramMaxSize 最大页大小
*/
func (p *PageInfo) RoundPageInfoEx(paramDefaultSize int, paramMaxSize int) *PageInfo {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = paramDefaultSize
	}
	if p.PageSize > paramMaxSize {
		p.PageSize = paramMaxSize
	}
	return p
}

// SetPage 设置页信息
func (p *PageInfo) SetPage(paramPage int, paramPageSize int) *PageInfo {
	p.Page = paramPage
	p.PageSize = paramPageSize
	return p.RoundPageInfo()
}

func (p *PageInfo) GetPage() int {
	return p.Page
}

func (p *PageInfo) GetPageSize() int {
	return p.PageSize
}

// CalcPageOffset 用于mysql的计算偏移
func (p *PageInfo) CalcPageOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PageInfo) Offset() int {
	return p.CalcPageOffset()
}

// CalcLimit 用于mysql的计算limit
func (p *PageInfo) CalcLimit() int {
	return p.PageSize
}

func (p *PageInfo) Limit() int {
	return p.CalcLimit()
}

// CalcMaxPage 根据记录数和页的大小 计算最大页数 paramPageSize <= 0时 计算失败
func CalcMaxPage(paramCount int, paramPageSize int) (*BaseRet, int) {
	r := BaseRet{}
	pageCnt := int(0)
	for range [1]int{} {
		if paramPageSize <= 0 {
			r.SetError(ERR_FAIL, fmt.Sprintf("paramPageSize = %d <= 0,  页数要是大于0的整数", paramPageSize))
			break
		}

		if paramCount <= 0 {
			break
		}
		p := paramCount % paramPageSize
		pageCnt = (paramCount - p) / paramPageSize
		if p > 0 {
			pageCnt++
		}
	}
	return &r, pageCnt
}
