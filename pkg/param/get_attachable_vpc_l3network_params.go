// Copyright (c) ZStack.io, Inc.

package param

// GetAttachableVpcL3NetworkDetailParam GetAttachableVpcL3Network detail param
type GetAttachableVpcL3NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetAttachableVpcL3NetworkParam GetAttachableVpcL3Network request param
type GetAttachableVpcL3NetworkParam struct {
	BaseParam
	Params GetAttachableVpcL3NetworkDetailParam `json:"params"`
}
