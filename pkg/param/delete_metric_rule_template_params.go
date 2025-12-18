// Copyright (c) ZStack.io, Inc.

package param

// DeleteMetricRuleTemplateDetailParam DeleteMetricRuleTemplate detail param
type DeleteMetricRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMetricRuleTemplateParam DeleteMetricRuleTemplate request param
type DeleteMetricRuleTemplateParam struct {
	BaseParam
	Params DeleteMetricRuleTemplateDetailParam `json:"params"`
}
