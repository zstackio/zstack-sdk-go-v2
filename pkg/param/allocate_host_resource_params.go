// Copyright (c) ZStack.io, Inc.

package param

// AllocateHostResourceDetailParam AllocateHostResource detail param
type AllocateHostResourceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Strategy string `json:"strategy" validate:"required"`
	Scene string `json:"scene" validate:"required"`
	Vcpu int `json:"vcpu" validate:"required"`
	MemSize int64 `json:"memSize,omitempty"`
}

// AllocateHostResourceParam AllocateHostResource request param
type AllocateHostResourceParam struct {
	BaseParam
	Params AllocateHostResourceDetailParam `json:"params"`
}
