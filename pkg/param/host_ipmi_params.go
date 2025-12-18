// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostIpmiDetailParam UpdateHostIpmi详细参数
type UpdateHostIpmiDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"ipmiAddress,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest string `json:"ipmiPassword,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
}

// UpdateHostIpmiParam UpdateHostIpmi请求参数
type UpdateHostIpmiParam struct {
	BaseParam
	Params UpdateHostIpmiDetailParam `json:"params"` // 详细参数
}

