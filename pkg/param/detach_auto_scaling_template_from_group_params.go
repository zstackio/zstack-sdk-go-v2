// Copyright (c) ZStack.io, Inc.

package param

// DetachAutoScalingTemplateFromGroupDetailParam DetachAutoScalingTemplateFromGroup detail param
type DetachAutoScalingTemplateFromGroupDetailParam struct {
	TemplateUuid string `json:"templateUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// DetachAutoScalingTemplateFromGroupParam DetachAutoScalingTemplateFromGroup request param
type DetachAutoScalingTemplateFromGroupParam struct {
	BaseParam
	Params DetachAutoScalingTemplateFromGroupDetailParam `json:"params"`
}
