// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmStaticIpDetailParam DeleteVmStaticIp detail param
type DeleteVmStaticIpDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StaticIp string `json:"staticIp,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmStaticIpParam DeleteVmStaticIp request param
type DeleteVmStaticIpParam struct {
	BaseParam
	Params DeleteVmStaticIpDetailParam `json:"params"`
}
