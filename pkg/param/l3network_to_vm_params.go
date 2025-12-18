// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworkToVmDetailParam AttachL3NetworkToVm详细参数
type AttachL3NetworkToVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"staticIp,omitempty"`
	rest string `json:"driverType,omitempty"`
	rest string `json:"customMac,omitempty"`
	rest string `json:"vmNicParams,omitempty"`
}

// AttachL3NetworkToVmParam AttachL3NetworkToVm请求参数
type AttachL3NetworkToVmParam struct {
	BaseParam
	Params AttachL3NetworkToVmDetailParam `json:"params"` // 详细参数
}

