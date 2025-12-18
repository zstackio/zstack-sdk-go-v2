// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmPasswordDetailParam ChangeVmPassword详细参数
type ChangeVmPasswordDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"account" validate:"required"` // 必填
}

// ChangeVmPasswordParam ChangeVmPassword请求参数
type ChangeVmPasswordParam struct {
	BaseParam
	Params ChangeVmPasswordDetailParam `json:"params"` // 详细参数
}

