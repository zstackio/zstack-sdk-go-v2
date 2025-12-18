// Copyright (c) ZStack.io, Inc.

package param

// AddActionToAlarmDetailParam AddActionToAlarm详细参数
type AddActionToAlarmDetailParam struct {
	rest string `json:"alarmUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
	rest string `json:"actionType" validate:"required"` // 必填
}

// AddActionToAlarmParam AddActionToAlarm请求参数
type AddActionToAlarmParam struct {
	BaseParam
	Params AddActionToAlarmDetailParam `json:"params"` // 详细参数
}

