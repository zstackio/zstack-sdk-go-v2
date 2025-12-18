// Copyright (c) ZStack.io, Inc.

package param

// ChangeAlarmStateDetailParam ChangeAlarmState详细参数
type ChangeAlarmStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeAlarmStateParam ChangeAlarmState请求参数
type ChangeAlarmStateParam struct {
	BaseParam
	Params ChangeAlarmStateDetailParam `json:"params"` // 详细参数
}

