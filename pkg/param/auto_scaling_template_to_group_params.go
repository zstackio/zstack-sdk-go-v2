// Copyright (c) ZStack.io, Inc.

package param

// AttachAutoScalingTemplateToGroupDetailParam AttachAutoScalingTemplateToGroup详细参数
type AttachAutoScalingTemplateToGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// AttachAutoScalingTemplateToGroupParam AttachAutoScalingTemplateToGroup请求参数
type AttachAutoScalingTemplateToGroupParam struct {
	BaseParam
	Params AttachAutoScalingTemplateToGroupDetailParam `json:"params"` // 详细参数
}

