// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2InstancePasswordDetailParam ChangeBareMetal2InstancePassword详细参数
type ChangeBareMetal2InstancePasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
}

// ChangeBareMetal2InstancePasswordParam ChangeBareMetal2InstancePassword请求参数
type ChangeBareMetal2InstancePasswordParam struct {
	BaseParam
	Params ChangeBareMetal2InstancePasswordDetailParam `json:"params"` // 详细参数
}

