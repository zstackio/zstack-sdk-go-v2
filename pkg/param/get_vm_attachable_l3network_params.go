// Copyright (c) ZStack.io, Inc.

package param

// GetVmAttachableL3NetworkDetailParam GetVmAttachableL3Network detail param
type GetVmAttachableL3NetworkDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetVmAttachableL3NetworkParam GetVmAttachableL3Network request param
type GetVmAttachableL3NetworkParam struct {
	BaseParam
	Params GetVmAttachableL3NetworkDetailParam `json:"params"`
}
