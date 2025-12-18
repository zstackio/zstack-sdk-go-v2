// Copyright (c) ZStack.io, Inc.

package param

// UpdateActiveAlarmTemplateDetailParam UpdateActiveAlarmTemplate详细参数
type UpdateActiveAlarmTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"alarmName,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest string `json:"labels,omitempty"`
}

// UpdateActiveAlarmTemplateParam UpdateActiveAlarmTemplate请求参数
type UpdateActiveAlarmTemplateParam struct {
	BaseParam
	Params UpdateActiveAlarmTemplateDetailParam `json:"params"` // 详细参数
}

