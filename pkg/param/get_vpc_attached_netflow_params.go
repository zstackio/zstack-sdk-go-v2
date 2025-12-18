// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedNetflowDetailParam GetVpcAttachedNetflow detail param
type GetVpcAttachedNetflowDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedNetflowParam GetVpcAttachedNetflow request param
type GetVpcAttachedNetflowParam struct {
	BaseParam
	Params GetVpcAttachedNetflowDetailParam `json:"params"`
}
