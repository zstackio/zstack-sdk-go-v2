// Copyright (c) ZStack.io, Inc.

package param

// GetHostResourceAllocationDetailParam GetHostResourceAllocation detail param
type GetHostResourceAllocationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Strategy string `json:"strategy" validate:"required"`
	Scene string `json:"scene" validate:"required"`
	Vcpu int `json:"vcpu" validate:"required"`
	MemSize int64 `json:"memSize,omitempty"`
}

// GetHostResourceAllocationParam GetHostResourceAllocation request param
type GetHostResourceAllocationParam struct {
	BaseParam
	Params GetHostResourceAllocationDetailParam `json:"params"`
}
