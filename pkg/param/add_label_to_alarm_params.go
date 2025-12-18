// Copyright (c) ZStack.io, Inc.

package param

// AddLabelToAlarmDetailParam AddLabelToAlarm detail param
type AddLabelToAlarmDetailParam struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddLabelToAlarmParam AddLabelToAlarm request param
type AddLabelToAlarmParam struct {
	BaseParam
	Params AddLabelToAlarmDetailParam `json:"params"`
}
