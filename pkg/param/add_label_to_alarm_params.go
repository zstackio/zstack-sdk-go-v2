// Copyright (c) ZStack.io, Inc.

package param

// AddLabelToAlarmDetailParam AddLabelToAlarm详细参数
type AddLabelToAlarmDetailParam struct {
	rest string `json:"alarmUuid" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
	rest string `json:"operator" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddLabelToAlarmParam AddLabelToAlarm请求参数
type AddLabelToAlarmParam struct {
	BaseParam
	Params AddLabelToAlarmDetailParam `json:"params"` // 详细参数
}

