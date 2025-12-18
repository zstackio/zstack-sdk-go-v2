// Copyright (c) ZStack.io, Inc.

package param

// AddActionToAlarmDetailParam AddActionToAlarm detail param
type AddActionToAlarmDetailParam struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
	ActionType string `json:"actionType" validate:"required"`
}

// AddActionToAlarmParam AddActionToAlarm request param
type AddActionToAlarmParam struct {
	BaseParam
	Params AddActionToAlarmDetailParam `json:"params"`
}
