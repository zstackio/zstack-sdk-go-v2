// Copyright (c) ZStack.io, Inc.

package param

// DestroyVmInstanceDetailParam DestroyVmInstance detail param
type DestroyVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DestroyVmInstanceParam DestroyVmInstance request param
type DestroyVmInstanceParam struct {
	BaseParam
	Params DestroyVmInstanceDetailParam `json:"params"`
}
