// Copyright (c) ZStack.io, Inc.

package param

// GetVmNicAttachedNetworkServiceDetailParam GetVmNicAttachedNetworkService detail param
type GetVmNicAttachedNetworkServiceDetailParam struct {
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
}

// GetVmNicAttachedNetworkServiceParam GetVmNicAttachedNetworkService request param
type GetVmNicAttachedNetworkServiceParam struct {
	BaseParam
	Params GetVmNicAttachedNetworkServiceDetailParam `json:"params"`
}
