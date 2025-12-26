// Copyright (c) ZStack.io, Inc.

package param

// UpdateMetricRuleTemplateDetailParam UpdateMetricRuleTemplate detail param
type UpdateMetricRuleTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
}

// UpdateMetricRuleTemplateParam UpdateMetricRuleTemplate request param
type UpdateMetricRuleTemplateParam struct {
	BaseParam
	Params UpdateMetricRuleTemplateDetailParam `json:"params"`
}
