// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedIpsecDetailParam GetVpcAttachedIpsec detail param
type GetVpcAttachedIpsecDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedIpsecParam GetVpcAttachedIpsec request param
type GetVpcAttachedIpsecParam struct {
	BaseParam
	Params GetVpcAttachedIpsecDetailParam `json:"params"`
}
