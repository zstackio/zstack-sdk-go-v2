// Copyright (c) ZStack.io, Inc.

package param

// ChangeActiveAlarmStateDetailParam ChangeActiveAlarmState详细参数
type ChangeActiveAlarmStateDetailParam struct {
	rest string `json:"namespace" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeActiveAlarmStateParam ChangeActiveAlarmState请求参数
type ChangeActiveAlarmStateParam struct {
	BaseParam
	Params ChangeActiveAlarmStateDetailParam `json:"params"` // 详细参数
}

