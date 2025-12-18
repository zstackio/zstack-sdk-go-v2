// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedEipDetailParam GetVpcAttachedEip详细参数
type GetVpcAttachedEipDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedEipParam GetVpcAttachedEip请求参数
type GetVpcAttachedEipParam struct {
	BaseParam
	Params GetVpcAttachedEipDetailParam `json:"params"` // 详细参数
}

