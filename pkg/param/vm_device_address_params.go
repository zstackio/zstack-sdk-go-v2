// Copyright (c) ZStack.io, Inc.

package param

// GetVmDeviceAddressDetailParam GetVmDeviceAddress详细参数
type GetVmDeviceAddressDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"resourceTypes" validate:"required"` // 必填
}

// GetVmDeviceAddressParam GetVmDeviceAddress请求参数
type GetVmDeviceAddressParam struct {
	BaseParam
	Params GetVmDeviceAddressDetailParam `json:"params"` // 详细参数
}

