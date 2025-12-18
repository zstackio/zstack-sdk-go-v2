// Copyright (c) ZStack.io, Inc.

package param

// GetVpcIPsecLogDetailParam GetVpcIPsecLog详细参数
type GetVpcIPsecLogDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"lines,omitempty"`
}

// GetVpcIPsecLogParam GetVpcIPsecLog请求参数
type GetVpcIPsecLogParam struct {
	BaseParam
	Params GetVpcIPsecLogDetailParam `json:"params"` // 详细参数
}

