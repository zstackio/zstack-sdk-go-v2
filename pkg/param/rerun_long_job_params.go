// Copyright (c) ZStack.io, Inc.

package param

// RerunLongJobDetailParam RerunLongJob详细参数
type RerunLongJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RerunLongJobParam RerunLongJob请求参数
type RerunLongJobParam struct {
	BaseParam
	Params RerunLongJobDetailParam `json:"params"` // 详细参数
}

