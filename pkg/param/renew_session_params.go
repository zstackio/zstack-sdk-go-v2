// Copyright (c) ZStack.io, Inc.

package param

// RenewSessionDetailParam RenewSession详细参数
type RenewSessionDetailParam struct {
	rest string `json:"sessionUuid" validate:"required"` // 必填
	rest int64 `json:"duration,omitempty"`
}

// RenewSessionParam RenewSession请求参数
type RenewSessionParam struct {
	BaseParam
	Params RenewSessionDetailParam `json:"params"` // 详细参数
}

