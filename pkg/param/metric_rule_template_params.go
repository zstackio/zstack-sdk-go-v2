// Copyright (c) ZStack.io, Inc.

package param

// UpdateMetricRuleTemplateDetailParam UpdateMetricRuleTemplate详细参数
type UpdateMetricRuleTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest []interface{} `json:"labels,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest bool `json:"enableRecovery,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
}

// UpdateMetricRuleTemplateParam UpdateMetricRuleTemplate请求参数
type UpdateMetricRuleTemplateParam struct {
	BaseParam
	Params UpdateMetricRuleTemplateDetailParam `json:"params"` // 详细参数
}

