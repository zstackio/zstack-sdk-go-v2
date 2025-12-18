// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmLabelDetailParam UpdateAlarmLabel详细参数
type UpdateAlarmLabelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"key" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
	rest string `json:"operator" validate:"required"` // 必填
}

// UpdateAlarmLabelParam UpdateAlarmLabel请求参数
type UpdateAlarmLabelParam struct {
	BaseParam
	Params UpdateAlarmLabelDetailParam `json:"params"` // 详细参数
}

