// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworkToVmNicDetailParam AttachL3NetworkToVmNic detail param
type AttachL3NetworkToVmNicDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StaticIp string `json:"staticIp,omitempty"`
}

// AttachL3NetworkToVmNicParam AttachL3NetworkToVmNic request param
type AttachL3NetworkToVmNicParam struct {
	BaseParam
	Params AttachL3NetworkToVmNicDetailParam `json:"params"`
}
