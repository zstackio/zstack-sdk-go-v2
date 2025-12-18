// Copyright (c) ZStack.io, Inc.

package param

// DetachAutoScalingTemplateFromGroupDetailParam DetachAutoScalingTemplateFromGroup详细参数
type DetachAutoScalingTemplateFromGroupDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// DetachAutoScalingTemplateFromGroupParam DetachAutoScalingTemplateFromGroup请求参数
type DetachAutoScalingTemplateFromGroupParam struct {
	BaseParam
	Params DetachAutoScalingTemplateFromGroupDetailParam `json:"params"` // 详细参数
}

