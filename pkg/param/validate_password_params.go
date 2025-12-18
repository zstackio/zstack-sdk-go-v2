// Copyright (c) ZStack.io, Inc.

package param

// ValidatePasswordDetailParam ValidatePassword详细参数
type ValidatePasswordDetailParam struct {
	rest string `json:"loginName" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
}

// ValidatePasswordParam ValidatePassword请求参数
type ValidatePasswordParam struct {
	BaseParam
	Params ValidatePasswordDetailParam `json:"params"` // 详细参数
}

