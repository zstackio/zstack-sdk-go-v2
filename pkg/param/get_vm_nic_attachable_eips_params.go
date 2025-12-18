// Copyright (c) ZStack.io, Inc.

package param

// GetVmNicAttachableEipsDetailParam GetVmNicAttachableEips detail param
type GetVmNicAttachableEipsDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
	IpVersion int `json:"ipVersion,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVmNicAttachableEipsParam GetVmNicAttachableEips request param
type GetVmNicAttachableEipsParam struct {
	BaseParam
	Params GetVmNicAttachableEipsDetailParam `json:"params"`
}
