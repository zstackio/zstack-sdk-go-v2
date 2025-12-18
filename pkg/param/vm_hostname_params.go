// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmHostnameDetailParam DeleteVmHostname详细参数
type DeleteVmHostnameDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmHostnameParam DeleteVmHostname请求参数
type DeleteVmHostnameParam struct {
	BaseParam
	Params DeleteVmHostnameDetailParam `json:"params"` // 详细参数
}

