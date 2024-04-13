package xm

import (
	"strconv"
)

const (
	ERR_OK   = 0
	ERR_FAIL = 1
)

type BaseRet struct {
	ret int
	msg string
}

// CommonRet 通用返回结构
// ret = 0 的时候，表示成功，其他值表示失败
type CommonRet[T any] struct {
	BaseRet
	data *T
}

type IRet interface {
	GetRet() int
	GetMsg() string
	IsOK() bool
	IsNotOK() bool
}

// CommonRetData 通用返回数据结构
type CommonRetData[T any] struct {
	// 返回码
	Ret int `json:"ret"`
	// 返回错误信息
	Msg string `json:"msg"`
	// 返回错误数据
	Data *T `json:"data"`
}

// SetRet 设置返回码
func (r *CommonRet[T]) SetRet(paramRet int) *CommonRet[T] {
	r.ret = paramRet
	return r
}

// GetRet 取返回码
func (r *CommonRet[T]) GetRet() int {
	return r.ret
}

// ToData 返回数据结构
func (r *CommonRet[T]) ToData() *CommonRetData[T] {
	return &CommonRetData[T]{
		Ret:  r.ret,
		Msg:  r.msg,
		Data: r.data,
	}
}

// GetRetStr 取返回码字符串值
func (r *CommonRet[T]) GetRetStr() string {
	return strconv.Itoa(r.ret)
}

// SetMsg 设置错误信息
func (r *CommonRet[T]) SetMsg(paramMsg string) *CommonRet[T] {
	r.msg = paramMsg
	return r
}

// GetMsg 取错误信息
func (r *CommonRet[T]) GetMsg() string {
	return r.msg
}

// SetError 设置错误
func (r *CommonRet[T]) SetError(paramRet int, paramMsg string) *CommonRet[T] {
	r.ret = paramRet
	r.msg = paramMsg
	return r
}

// GetError 取Error信息
func (r *CommonRet[T]) GetError() (int, string) {
	return r.ret, r.msg
}

// SetOK 设置成功
func (r *CommonRet[T]) SetOK(paramData *T) *CommonRet[T] {
	r.ret = ERR_OK
	r.msg = ""
	r.data = paramData
	return r
}

// SetData 设置返回数据
func (r *CommonRet[T]) SetData(paramData *T) *CommonRet[T] {
	r.data = paramData
	return r
}

// 实现error接口
func (r *CommonRet[T]) Error() string {
	return r.msg
}

// GetData 取返回数据
func (r *CommonRet[T]) GetData() *T {
	return r.data
}

// IsOK 判断是否成功
func (r *CommonRet[T]) IsOK() bool {
	return r.ret == ERR_OK
}

// IsNotOK 判断是否成功
func (r *CommonRet[T]) IsNotOK() bool {
	return !r.IsOK()
}

// Reset 重置数据
func (r *CommonRet[T]) Reset() {
	r.ret = ERR_OK
	r.msg = ""
	r.data = nil
}

// AssignFrom 从另一个ret赋值
func (r *CommonRet[T]) AssignFrom(paramR *CommonRet[T]) *CommonRet[T] {
	r.ret = paramR.ret
	r.msg = paramR.msg
	r.data = paramR.data
	return r
}

// ToRetNoData 返回无数据的一个整数的Common
func (r *CommonRet[T]) ToRetNoData() *CommonRet[int] {
	rrr := &CommonRet[int]{}
	rrr.ret = r.ret
	rrr.msg = r.msg
	return rrr
}

// ToBaseRet
func (r *CommonRet[T]) ToBaseRet() *BaseRet {
	rrr := &BaseRet{
		ret: r.ret,
		msg: r.msg,
	}
	return rrr
}

// AssignErrorFrom 复制错误信息
func (r *CommonRet[T]) AssignErrorFrom(paramR IRet) *CommonRet[T] {
	r.ret = paramR.GetRet()
	r.msg = paramR.GetMsg()
	return r
}

// SetRet 设置返回码
func (r *BaseRet) SetRet(paramRet int) *BaseRet {
	r.ret = paramRet
	return r
}

// GetRet 取返回码
func (r *BaseRet) GetRet() int {
	return r.ret
}

// GetRetStr 取返回码字符串值
func (r *BaseRet) GetRetStr() string {
	return strconv.Itoa(r.ret)
}

// SetMsg 设置错误信息
func (r *BaseRet) SetMsg(paramMsg string) *BaseRet {
	r.msg = paramMsg
	return r
}

// GetMsg 取错误信息
func (r *BaseRet) GetMsg() string {
	return r.msg
}

// SetError 设置错误
func (r *BaseRet) SetError(paramRet int, paramMsg string) *BaseRet {
	r.ret = paramRet
	r.msg = paramMsg
	return r
}

// GetError 取Error信息
func (r *BaseRet) GetError() (int, string) {
	return r.ret, r.msg
}

// SetOK 设置成功
func (r *BaseRet) SetOK() *BaseRet {
	r.ret = ERR_OK
	r.msg = ""
	return r
}

// 实现error接口
func (r *BaseRet) Error() string {
	return r.msg
}

// IsOK 判断是否成功
func (r *BaseRet) IsOK() bool {
	return r.ret == ERR_OK
}

// IsNotOK 判断是否成功
func (r *BaseRet) IsNotOK() bool {
	return !r.IsOK()
}

// Reset 重置数据
func (r *BaseRet) Reset() {
	r.ret = ERR_OK
	r.msg = ""
}

// AssignFrom 从另一个ret赋值
func (r *BaseRet) AssignFrom(paramR *BaseRet) *BaseRet {
	r.ret = paramR.ret
	r.msg = paramR.msg
	return r
}

// AssignErrorFrom 复制错误信息
func (r *BaseRet) AssignErrorFrom(paramR IRet) *BaseRet {
	r.ret = paramR.GetRet()
	r.msg = paramR.GetMsg()
	return r
}
