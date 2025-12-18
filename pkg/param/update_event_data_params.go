// Copyright (c) ZStack.io, Inc.

package param

// UpdateEventDataDetailParam UpdateEventData detail param
type UpdateEventDataDetailParam struct {
	DataUuid string `json:"dataUuid,omitempty"`
	DataStartTime int64 `json:"dataStartTime,omitempty"`
	DataEndTime int64 `json:"dataEndTime,omitempty"`
	UpdateMode string `json:"updateMode" validate:"required"`
	ReadStatus string `json:"readStatus,omitempty"`
}

// UpdateEventDataParam UpdateEventData request param
type UpdateEventDataParam struct {
	BaseParam
	Params UpdateEventDataDetailParam `json:"params"`
}
