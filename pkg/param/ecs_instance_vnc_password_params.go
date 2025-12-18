// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsInstanceVncPasswordDetailParam UpdateEcsInstanceVncPassword详细参数
type UpdateEcsInstanceVncPasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
}

// UpdateEcsInstanceVncPasswordParam UpdateEcsInstanceVncPassword请求参数
type UpdateEcsInstanceVncPasswordParam struct {
	BaseParam
	Params UpdateEcsInstanceVncPasswordDetailParam `json:"params"` // 详细参数
}

