// Copyright (c) ZStack.io, Inc.

package param

// ValidateSessionDetailParam ValidateSession详细参数
type ValidateSessionDetailParam struct {
	rest string `json:"sessionUuid" validate:"required"` // 必填
}

// ValidateSessionParam ValidateSession请求参数
type ValidateSessionParam struct {
	BaseParam
	Params ValidateSessionDetailParam `json:"params"` // 详细参数
}

