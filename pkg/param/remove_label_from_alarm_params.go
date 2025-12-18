// Copyright (c) ZStack.io, Inc.

package param

// RemoveLabelFromAlarmDetailParam RemoveLabelFromAlarm详细参数
type RemoveLabelFromAlarmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveLabelFromAlarmParam RemoveLabelFromAlarm请求参数
type RemoveLabelFromAlarmParam struct {
	BaseParam
	Params RemoveLabelFromAlarmDetailParam `json:"params"` // 详细参数
}

