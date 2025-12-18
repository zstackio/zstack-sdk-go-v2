// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmDataDetailParam UpdateAlarmData详细参数
type UpdateAlarmDataDetailParam struct {
	rest string `json:"dataUuid,omitempty"`
	rest int64 `json:"dataStartTime,omitempty"`
	rest int64 `json:"dataEndTime,omitempty"`
	rest string `json:"updateMode" validate:"required"` // 必填
	rest string `json:"readStatus,omitempty"`
}

// UpdateAlarmDataParam UpdateAlarmData请求参数
type UpdateAlarmDataParam struct {
	BaseParam
	Params UpdateAlarmDataDetailParam `json:"params"` // 详细参数
}

