// Copyright (c) ZStack.io, Inc.

package param

// AttachMdevDeviceToVmDetailParam AttachMdevDeviceToVm详细参数
type AttachMdevDeviceToVmDetailParam struct {
	rest string `json:"mdevDeviceUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// AttachMdevDeviceToVmParam AttachMdevDeviceToVm请求参数
type AttachMdevDeviceToVmParam struct {
	BaseParam
	Params AttachMdevDeviceToVmDetailParam `json:"params"` // 详细参数
}

