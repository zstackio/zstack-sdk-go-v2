// Copyright (c) ZStack.io, Inc.

package param

// AttachAutoScalingTemplateToGroupDetailParam AttachAutoScalingTemplateToGroup detail param
type AttachAutoScalingTemplateToGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// AttachAutoScalingTemplateToGroupParam AttachAutoScalingTemplateToGroup request param
type AttachAutoScalingTemplateToGroupParam struct {
	BaseParam
	Params AttachAutoScalingTemplateToGroupDetailParam `json:"params"`
}
