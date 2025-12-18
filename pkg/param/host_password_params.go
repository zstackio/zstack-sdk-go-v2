// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostPasswordDetailParam ChangeHostPassword详细参数
type ChangeHostPasswordDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
}

// ChangeHostPasswordParam ChangeHostPassword请求参数
type ChangeHostPasswordParam struct {
	BaseParam
	Params ChangeHostPasswordDetailParam `json:"params"` // 详细参数
}

