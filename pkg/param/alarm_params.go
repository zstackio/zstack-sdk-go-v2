// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmDetailParam UpdateAlarm详细参数
type UpdateAlarmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"comparisonOperator,omitempty"`
	rest int `json:"period,omitempty"`
	rest float64 `json:"threshold,omitempty"`
	rest int `json:"repeatInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest bool `json:"enableRecovery,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
	rest []interface{} `json:"actions,omitempty"`
}

// UpdateAlarmParam UpdateAlarm请求参数
type UpdateAlarmParam struct {
	BaseParam
	Params UpdateAlarmDetailParam `json:"params"` // 详细参数
}

