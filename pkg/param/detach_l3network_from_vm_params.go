// Copyright (c) ZStack.io, Inc.

package param

// DetachL3NetworkFromVmDetailParam DetachL3NetworkFromVm detail param
type DetachL3NetworkFromVmDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
}

// DetachL3NetworkFromVmParam DetachL3NetworkFromVm request param
type DetachL3NetworkFromVmParam struct {
	BaseParam
	Params DetachL3NetworkFromVmDetailParam `json:"params"`
}
