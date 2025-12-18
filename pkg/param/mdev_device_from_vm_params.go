// Copyright (c) ZStack.io, Inc.

package param

// DetachMdevDeviceFromVmDetailParam DetachMdevDeviceFromVm详细参数
type DetachMdevDeviceFromVmDetailParam struct {
	rest string `json:"mdevDeviceUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// DetachMdevDeviceFromVmParam DetachMdevDeviceFromVm请求参数
type DetachMdevDeviceFromVmParam struct {
	BaseParam
	Params DetachMdevDeviceFromVmDetailParam `json:"params"` // 详细参数
}

