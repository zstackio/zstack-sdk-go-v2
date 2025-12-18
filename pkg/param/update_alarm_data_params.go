// Copyright (c) ZStack.io, Inc.

package param

// UpdateAlarmDataDetailParam UpdateAlarmData detail param
type UpdateAlarmDataDetailParam struct {
	DataUuid string `json:"dataUuid,omitempty"`
	DataStartTime int64 `json:"dataStartTime,omitempty"`
	DataEndTime int64 `json:"dataEndTime,omitempty"`
	UpdateMode string `json:"updateMode" validate:"required"`
	ReadStatus string `json:"readStatus,omitempty"`
}

// UpdateAlarmDataParam UpdateAlarmData request param
type UpdateAlarmDataParam struct {
	BaseParam
	Params UpdateAlarmDataDetailParam `json:"params"`
}
