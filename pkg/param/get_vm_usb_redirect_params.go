// Copyright (c) ZStack.io, Inc.

package param

// GetVmUsbRedirectDetailParam GetVmUsbRedirect detail param
type GetVmUsbRedirectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmUsbRedirectParam GetVmUsbRedirect request param
type GetVmUsbRedirectParam struct {
	BaseParam
	Params GetVmUsbRedirectDetailParam `json:"params"`
}
