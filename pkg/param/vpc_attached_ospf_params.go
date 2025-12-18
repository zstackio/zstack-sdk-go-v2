// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedOspfDetailParam GetVpcAttachedOspf详细参数
type GetVpcAttachedOspfDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedOspfParam GetVpcAttachedOspf请求参数
type GetVpcAttachedOspfParam struct {
	BaseParam
	Params GetVpcAttachedOspfDetailParam `json:"params"` // 详细参数
}

