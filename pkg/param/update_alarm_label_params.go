// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmLabelDetailParam UpdateAlarmLabel detail param
type UpdateAlarmLabelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Key string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
	Operator string `json:"operator" validate:"required"`
}

// UpdateAlarmLabelParam UpdateAlarmLabel request param
type UpdateAlarmLabelParam struct {
	BaseParam
	Params UpdateAlarmLabelDetailParam `json:"params"`
}
