// Copyright (c) ZStack.io, Inc.

package param

// CleanLongJobDetailParam CleanLongJob详细参数
type CleanLongJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// CleanLongJobParam CleanLongJob请求参数
type CleanLongJobParam struct {
	BaseParam
	Params CleanLongJobDetailParam `json:"params"` // 详细参数
}

