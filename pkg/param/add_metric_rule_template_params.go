// Copyright (c) ZStack.io, Inc.

package param

// AddMetricRuleTemplateDetailParam AddMetricRuleTemplate详细参数
type AddMetricRuleTemplateDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"monitorTemplateUuid" validate:"required"` // 必填
	rest string `json:"comparisonOperator" validate:"required"` // 必填
	rest int `json:"period,omitempty"`
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"metricName" validate:"required"` // 必填
	rest float64 `json:"threshold" validate:"required"` // 必填
	rest int `json:"repeatInterval,omitempty"`
	rest []interface{} `json:"labels,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest bool `json:"enableRecovery,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddMetricRuleTemplateParam AddMetricRuleTemplate请求参数
type AddMetricRuleTemplateParam struct {
	BaseParam
	Params AddMetricRuleTemplateDetailParam `json:"params"` // 详细参数
}

