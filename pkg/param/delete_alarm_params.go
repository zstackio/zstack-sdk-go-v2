// Copyright (c) ZStack.io, Inc.

package param

// DeleteAlarmDetailParam DeleteAlarm detail param
type DeleteAlarmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAlarmParam DeleteAlarm request param
type DeleteAlarmParam struct {
	BaseParam
	Params DeleteAlarmDetailParam `json:"params"`
}
