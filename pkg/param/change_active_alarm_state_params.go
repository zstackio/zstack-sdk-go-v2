// Copyright (c) ZStack.io, Inc.

package param

// ChangeActiveAlarmStateDetailParam ChangeActiveAlarmState detail param
type ChangeActiveAlarmStateDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeActiveAlarmStateParam ChangeActiveAlarmState request param
type ChangeActiveAlarmStateParam struct {
	BaseParam
	Params ChangeActiveAlarmStateDetailParam `json:"params"`
}
