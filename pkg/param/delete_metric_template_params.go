// Copyright (c) ZStack.io, Inc.

package param

// DeleteMetricTemplateDetailParam DeleteMetricTemplate detail param
type DeleteMetricTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMetricTemplateParam DeleteMetricTemplate request param
type DeleteMetricTemplateParam struct {
	BaseParam
	Params DeleteMetricTemplateDetailParam `json:"params"`
}
