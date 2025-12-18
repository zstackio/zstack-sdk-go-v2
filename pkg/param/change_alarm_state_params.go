// Copyright (c) ZStack.io, Inc.

package param

// ChangeAlarmStateDetailParam ChangeAlarmState detail param
type ChangeAlarmStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAlarmStateParam ChangeAlarmState request param
type ChangeAlarmStateParam struct {
	BaseParam
	Params ChangeAlarmStateDetailParam `json:"params"`
}
