// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedVipDetailParam GetVpcAttachedVip详细参数
type GetVpcAttachedVipDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedVipParam GetVpcAttachedVip请求参数
type GetVpcAttachedVipParam struct {
	BaseParam
	Params GetVpcAttachedVipDetailParam `json:"params"` // 详细参数
}

