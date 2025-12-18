// Copyright (c) ZStack.io, Inc.

package param

// SetVmUsbRedirectDetailParam SetVmUsbRedirect详细参数
type SetVmUsbRedirectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetVmUsbRedirectParam SetVmUsbRedirect请求参数
type SetVmUsbRedirectParam struct {
	BaseParam
	Params SetVmUsbRedirectDetailParam `json:"params"` // 详细参数
}

