// Copyright (c) ZStack.io, Inc.

package param

// ReimageVmInstanceDetailParam ReimageVmInstance detail param
type ReimageVmInstanceDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ReimageVmInstanceParam ReimageVmInstance request param
type ReimageVmInstanceParam struct {
	BaseParam
	Params ReimageVmInstanceDetailParam `json:"params"`
}
