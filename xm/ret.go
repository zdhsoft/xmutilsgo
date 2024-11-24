package xm

import (
	"fmt"
	"strconv"
)

const (
	ERR_OK   = 0
	ERR_FAIL = 1
)

type IRet interface {
	GetRet() int
	GetMsg() string
	IsOK() bool
	IsNotOK() bool
}

type BaseRet struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

// SetRet 设置返回码
func (r *BaseRet) SetRet(paramRet int) *BaseRet {
	r.Ret = paramRet
	return r
}

// GetRet 取返回码
func (r *BaseRet) GetRet() int {
	return r.Ret
}

// GetRetStr 取返回码字符串值
func (r *BaseRet) GetRetStr() string {
	return strconv.Itoa(r.Ret)
}

// SetMsg 设置错误信息
func (r *BaseRet) SetMsg(paramMsg string) *BaseRet {
	r.Msg = paramMsg
	return r
}

// GetMsg 取错误信息
func (r *BaseRet) GetMsg() string {
	return r.Msg
}

// SetError 设置错误
func (r *BaseRet) SetError(paramRet int, paramMsg string) *BaseRet {
	r.Ret = paramRet
	r.Msg = paramMsg
	return r
}

// GetError 取Error信息
func (r *BaseRet) GetError() (int, string) {
	return r.Ret, r.Msg
}

// SetOK 设置成功
func (r *BaseRet) SetOK() *BaseRet {
	r.Ret = ERR_OK
	r.Msg = ""
	return r
}

// 实现error接口
func (r *BaseRet) Error() string {
	return fmt.Sprintf("[ret:%d]%s", r.Ret, r.Msg)
}

// IsOK 判断是否成功
func (r *BaseRet) IsOK() bool {
	return r.Ret == ERR_OK
}

// IsNotOK 判断是否成功
func (r *BaseRet) IsNotOK() bool {
	return !r.IsOK()
}

// Reset 重置数据
func (r *BaseRet) Reset() {
	r.Ret = ERR_OK
	r.Msg = ""
}

// AssignFrom 从另一个ret赋值
func (r *BaseRet) AssignFrom(paramR *BaseRet) *BaseRet {
	r.Ret = paramR.Ret
	r.Msg = paramR.Msg
	return r
}

// AssignErrorFrom 从另一个错误ret复制
func (r *BaseRet) AssignErrorFrom(paramR IRet) {
	r.Ret = paramR.GetRet()
	r.Msg = paramR.GetMsg()
}

/*
NewBaseRet 创建一个默认的BaseRet
*/
func NewBaseRet() *BaseRet {
	return new(BaseRet)
}

/*
NewBaseRetInit 创建BaseRet并初始化
*/
func NewBaseRetInit(paramRet int, paramMsg string) *BaseRet {
	r := &BaseRet{
		Ret: paramRet,
		Msg: paramMsg,
	}
	return r
}

/*
NewBaseRetError 创建BaseRet并初始化
*/
func NewBaseRetError(paramRet int, paramMsg string) error {
	r := &BaseRet{
		Ret: paramRet,
		Msg: paramMsg,
	}
	return r
}

// Error2Ret 将error转换为BaseRet
func Error2Ret(err error) *BaseRet {
	if err == nil {
		return nil
	}
	r := err.(*BaseRet)
	return r
}
