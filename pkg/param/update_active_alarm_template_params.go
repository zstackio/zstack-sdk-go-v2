// Copyright (c) ZStack.io, Inc.

package param

// UpdateActiveAlarmTemplateDetailParam UpdateActiveAlarmTemplate detail param
type UpdateActiveAlarmTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AlarmName string `json:"alarmName,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
}

// UpdateActiveAlarmTemplateParam UpdateActiveAlarmTemplate request param
type UpdateActiveAlarmTemplateParam struct {
	BaseParam
	Params UpdateActiveAlarmTemplateDetailParam `json:"params"`
}
