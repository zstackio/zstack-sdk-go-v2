// Copyright (c) ZStack.io, Inc.

package param

// RemoveLabelFromAlarmDetailParam RemoveLabelFromAlarm detail param
type RemoveLabelFromAlarmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// RemoveLabelFromAlarmParam RemoveLabelFromAlarm request param
type RemoveLabelFromAlarmParam struct {
	BaseParam
	Params RemoveLabelFromAlarmDetailParam `json:"params"`
}
