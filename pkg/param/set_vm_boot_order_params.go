// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootOrderDetailParam SetVmBootOrder detail param
type SetVmBootOrderDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BootOrder []string `json:"bootOrder,omitempty"`
}

// SetVmBootOrderParam SetVmBootOrder request param
type SetVmBootOrderParam struct {
	BaseParam
	Params SetVmBootOrderDetailParam `json:"params"`
}
