// Copyright (c) ZStack.io, Inc.

package param

// AckAlarmDataDetailParam AckAlarmData detail param
type AckAlarmDataDetailParam struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	AlertDataUuid string `json:"alertDataUuid" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	AckPeriodSec int `json:"ackPeriodSec" validate:"required"`
}

// AckAlarmDataParam AckAlarmData request param
type AckAlarmDataParam struct {
	BaseParam
	Params AckAlarmDataDetailParam `json:"params"`
}
