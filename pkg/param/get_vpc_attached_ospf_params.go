// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedOspfDetailParam GetVpcAttachedOspf detail param
type GetVpcAttachedOspfDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedOspfParam GetVpcAttachedOspf request param
type GetVpcAttachedOspfParam struct {
	BaseParam
	Params GetVpcAttachedOspfDetailParam `json:"params"`
}
