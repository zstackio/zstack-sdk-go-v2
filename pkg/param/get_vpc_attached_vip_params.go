// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedVipDetailParam GetVpcAttachedVip detail param
type GetVpcAttachedVipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedVipParam GetVpcAttachedVip request param
type GetVpcAttachedVipParam struct {
	BaseParam
	Params GetVpcAttachedVipDetailParam `json:"params"`
}
