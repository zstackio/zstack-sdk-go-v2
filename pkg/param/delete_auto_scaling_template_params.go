// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingTemplateDetailParam DeleteAutoScalingTemplate detail param
type DeleteAutoScalingTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingTemplateParam DeleteAutoScalingTemplate request param
type DeleteAutoScalingTemplateParam struct {
	BaseParam
	Params DeleteAutoScalingTemplateDetailParam `json:"params"`
}
