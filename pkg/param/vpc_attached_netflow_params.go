// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedNetflowDetailParam GetVpcAttachedNetflow详细参数
type GetVpcAttachedNetflowDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedNetflowParam GetVpcAttachedNetflow请求参数
type GetVpcAttachedNetflowParam struct {
	BaseParam
	Params GetVpcAttachedNetflowDetailParam `json:"params"` // 详细参数
}

