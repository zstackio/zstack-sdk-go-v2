// Copyright (c) ZStack.io, Inc.

package param

// DeleteLogServerDetailParam DeleteLogServer详细参数
type DeleteLogServerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteLogServerParam DeleteLogServer请求参数
type DeleteLogServerParam struct {
	BaseParam
	Params DeleteLogServerDetailParam `json:"params"` // 详细参数
}

