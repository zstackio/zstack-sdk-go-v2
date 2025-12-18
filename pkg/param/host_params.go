// Copyright (c) ZStack.io, Inc.

package param

// DeleteHostDetailParam DeleteHost详细参数
type DeleteHostDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHostParam DeleteHost请求参数
type DeleteHostParam struct {
	BaseParam
	Params DeleteHostDetailParam `json:"params"` // 详细参数
}

