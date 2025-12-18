// Copyright (c) ZStack.io, Inc.

package param

// UpdateEventDataDetailParam UpdateEventData详细参数
type UpdateEventDataDetailParam struct {
	rest string `json:"dataUuid,omitempty"`
	rest int64 `json:"dataStartTime,omitempty"`
	rest int64 `json:"dataEndTime,omitempty"`
	rest string `json:"updateMode" validate:"required"` // 必填
	rest string `json:"readStatus,omitempty"`
}

// UpdateEventDataParam UpdateEventData请求参数
type UpdateEventDataParam struct {
	BaseParam
	Params UpdateEventDataDetailParam `json:"params"` // 详细参数
}

