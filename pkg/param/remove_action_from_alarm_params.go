// Copyright (c) ZStack.io, Inc.

package param

// RemoveActionFromAlarmDetailParam RemoveActionFromAlarm detail param
type RemoveActionFromAlarmDetailParam struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	ActionUuid string `json:"actionUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveActionFromAlarmParam RemoveActionFromAlarm request param
type RemoveActionFromAlarmParam struct {
	BaseParam
	Params RemoveActionFromAlarmDetailParam `json:"params"`
}
