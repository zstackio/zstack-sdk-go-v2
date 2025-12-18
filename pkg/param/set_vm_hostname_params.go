// Copyright (c) ZStack.io, Inc.

package param

// SetVmHostnameDetailParam SetVmHostname详细参数
type SetVmHostnameDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"hostname" validate:"required"` // 必填
}

// SetVmHostnameParam SetVmHostname请求参数
type SetVmHostnameParam struct {
	BaseParam
	Params SetVmHostnameDetailParam `json:"params"` // 详细参数
}

