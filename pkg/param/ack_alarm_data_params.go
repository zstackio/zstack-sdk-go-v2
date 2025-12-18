// Copyright (c) ZStack.io, Inc.

package param

// AckAlarmDataDetailParam AckAlarmData详细参数
type AckAlarmDataDetailParam struct {
	rest string `json:"alarmUuid" validate:"required"` // 必填
	rest string `json:"alertDataUuid" validate:"required"` // 必填
	rest string `json:"dataType" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest int `json:"ackPeriodSec" validate:"required"` // 必填
}

// AckAlarmDataParam AckAlarmData请求参数
type AckAlarmDataParam struct {
	BaseParam
	Params AckAlarmDataDetailParam `json:"params"` // 详细参数
}

