// Copyright (c) ZStack.io, Inc.

package param

// GetVmBootOrderDetailParam GetVmBootOrder detail param
type GetVmBootOrderDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmBootOrderParam GetVmBootOrder request param
type GetVmBootOrderParam struct {
	BaseParam
	Params GetVmBootOrderDetailParam `json:"params"`
}
