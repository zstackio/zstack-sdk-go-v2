// Copyright (c) ZStack.io, Inc.

package param

// DeleteSSOClientDetailParam DeleteSSOClient详细参数
type DeleteSSOClientDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSSOClientParam DeleteSSOClient请求参数
type DeleteSSOClientParam struct {
	BaseParam
	Params DeleteSSOClientDetailParam `json:"params"` // 详细参数
}

