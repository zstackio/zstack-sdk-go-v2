// Copyright (c) ZStack.io, Inc.

package param

// SetVmUsbRedirectDetailParam SetVmUsbRedirect detail param
type SetVmUsbRedirectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetVmUsbRedirectParam SetVmUsbRedirect request param
type SetVmUsbRedirectParam struct {
	BaseParam
	Params SetVmUsbRedirectDetailParam `json:"params"`
}
