// Copyright (c) ZStack.io, Inc.

package param

// GetEcsInstanceVncUrlDetailParam GetEcsInstanceVncUrl详细参数
type GetEcsInstanceVncUrlDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetEcsInstanceVncUrlParam GetEcsInstanceVncUrl请求参数
type GetEcsInstanceVncUrlParam struct {
	BaseParam
	Params GetEcsInstanceVncUrlDetailParam `json:"params"` // 详细参数
}

