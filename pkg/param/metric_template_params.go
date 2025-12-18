// Copyright (c) ZStack.io, Inc.

package param

// DeleteMetricTemplateDetailParam DeleteMetricTemplate详细参数
type DeleteMetricTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteMetricTemplateParam DeleteMetricTemplate请求参数
type DeleteMetricTemplateParam struct {
	BaseParam
	Params DeleteMetricTemplateDetailParam `json:"params"` // 详细参数
}

