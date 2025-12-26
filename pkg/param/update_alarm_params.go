// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmDetailParam UpdateAlarm detail param
type UpdateAlarmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ComparisonOperator string `json:"comparisonOperator,omitempty"`
	Period int `json:"period,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Actions []ActionParamParam `json:"actions,omitempty"`
}

// UpdateAlarmParam UpdateAlarm request param
type UpdateAlarmParam struct {
	BaseParam
	Params UpdateAlarmDetailParam `json:"params"`
}
