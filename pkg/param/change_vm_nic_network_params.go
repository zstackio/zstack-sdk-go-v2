// Copyright (c) ZStack.io, Inc.

package param

// ChangeVmNicNetworkDetailParam ChangeVmNicNetwork detail param
type ChangeVmNicNetworkDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	DestL3NetworkUuid string `json:"destL3NetworkUuid" validate:"required"`
	StaticIp string `json:"staticIp,omitempty"`
}

// ChangeVmNicNetworkParam ChangeVmNicNetwork request param
type ChangeVmNicNetworkParam struct {
	BaseParam
	Params ChangeVmNicNetworkDetailParam `json:"params"`
}
