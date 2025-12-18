// Copyright (c) ZStack.io, Inc.

package param

// RemoveActionFromAlarmDetailParam RemoveActionFromAlarm详细参数
type RemoveActionFromAlarmDetailParam struct {
	rest string `json:"alarmUuid" validate:"required"` // 必填
	rest string `json:"actionUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveActionFromAlarmParam RemoveActionFromAlarm请求参数
type RemoveActionFromAlarmParam struct {
	BaseParam
	Params RemoveActionFromAlarmDetailParam `json:"params"` // 详细参数
}

