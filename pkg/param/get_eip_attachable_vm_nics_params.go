// Copyright (c) ZStack.io, Inc.

package param

// GetEipAttachableVmNicsDetailParam GetEipAttachableVmNics detail param
type GetEipAttachableVmNicsDetailParam struct {
	EipUuid string `json:"eipUuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	VmName string `json:"vmName,omitempty"`
	NetworkServiceProvider string `json:"networkServiceProvider,omitempty"`
	AttachedToVm bool `json:"attachedToVm,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetEipAttachableVmNicsParam GetEipAttachableVmNics request param
type GetEipAttachableVmNicsParam struct {
	BaseParam
	Params GetEipAttachableVmNicsDetailParam `json:"params"`
}
