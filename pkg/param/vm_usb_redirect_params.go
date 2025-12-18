// Copyright (c) ZStack.io, Inc.

package param

// GetVmUsbRedirectDetailParam GetVmUsbRedirect详细参数
type GetVmUsbRedirectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmUsbRedirectParam GetVmUsbRedirect请求参数
type GetVmUsbRedirectParam struct {
	BaseParam
	Params GetVmUsbRedirectDetailParam `json:"params"` // 详细参数
}

