// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedEipDetailParam GetVpcAttachedEip detail param
type GetVpcAttachedEipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedEipParam GetVpcAttachedEip request param
type GetVpcAttachedEipParam struct {
	BaseParam
	Params GetVpcAttachedEipDetailParam `json:"params"`
}
