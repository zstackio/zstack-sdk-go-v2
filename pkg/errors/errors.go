// Copyright (c) ZStack.io, Inc.

package errors

import "fmt"

// Error 自定义错误类型
type Error string

func (e Error) Error() string {
	return string(e)
}

// 预定义错误常量
const (
	ErrNotFound    = Error("NotFoundError")    // 资源不存在
	ErrDuplicateId = Error("DuplicateIdError") // 重复ID错误
	ErrParameter   = Error("ParameterError")   // 参数错误
	ErrAuth        = Error("AuthError")        // 认证错误
	ErrPermission  = Error("PermissionError")  // 权限错误
	ErrInternal    = Error("InternalError")    // 内部错误
)

// Wrap 包装错误，添加上下文信息
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// Wrapf 格式化包装错误
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}

// ErrorCode API 错误码结构
type ErrorCode struct {
	Code        string     `json:"code"`        // 错误码
	Description string     `json:"description"` // 错误描述
	Details     string     `json:"details"`     // 详细信息
	Elaboration string     `json:"elaboration"` // 错误说明
	Cause       *ErrorCode `json:"cause"`       // 原因
}
