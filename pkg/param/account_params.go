// Copyright (c) ZStack.io, Inc.

package param

// DeleteAccountDetailParam DeleteAccount详细参数
type DeleteAccountDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAccountParam DeleteAccount请求参数
type DeleteAccountParam struct {
	BaseParam
	Params DeleteAccountDetailParam `json:"params"` // 详细参数
}

