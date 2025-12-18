// Copyright (c) ZStack.io, Inc.

package param

// AddMetricRuleTemplateDetailParam AddMetricRuleTemplate detail param
type AddMetricRuleTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid" validate:"required"`
	ComparisonOperator string `json:"comparisonOperator" validate:"required"`
	Period int `json:"period,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	Labels []interface{} `json:"labels,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddMetricRuleTemplateParam AddMetricRuleTemplate request param
type AddMetricRuleTemplateParam struct {
	BaseParam
	Params AddMetricRuleTemplateDetailParam `json:"params"`
}
