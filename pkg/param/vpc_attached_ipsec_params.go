// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedIpsecDetailParam GetVpcAttachedIpsec详细参数
type GetVpcAttachedIpsecDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedIpsecParam GetVpcAttachedIpsec请求参数
type GetVpcAttachedIpsecParam struct {
	BaseParam
	Params GetVpcAttachedIpsecDetailParam `json:"params"` // 详细参数
}

