// Copyright (c) ZStack.io, Inc.

package param

// AttachL3NetworkToVmDetailParam AttachL3NetworkToVm detail param
type AttachL3NetworkToVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StaticIp string `json:"staticIp,omitempty"`
	DriverType string `json:"driverType,omitempty"`
	CustomMac string `json:"customMac,omitempty"`
	VmNicParams string `json:"vmNicParams,omitempty"`
}

// AttachL3NetworkToVmParam AttachL3NetworkToVm request param
type AttachL3NetworkToVmParam struct {
	BaseParam
	Params AttachL3NetworkToVmDetailParam `json:"params"`
}
