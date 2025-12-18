// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingTemplateDetailParam DeleteAutoScalingTemplate详细参数
type DeleteAutoScalingTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingTemplateParam DeleteAutoScalingTemplate请求参数
type DeleteAutoScalingTemplateParam struct {
	BaseParam
	Params DeleteAutoScalingTemplateDetailParam `json:"params"` // 详细参数
}

