// Copyright (c) ZStack.io, Inc.

package param

// RemoveInstanceFromMonitorGroupDetailParam RemoveInstanceFromMonitorGroup详细参数
type RemoveInstanceFromMonitorGroupDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"instanceUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveInstanceFromMonitorGroupParam RemoveInstanceFromMonitorGroup请求参数
type RemoveInstanceFromMonitorGroupParam struct {
	BaseParam
	Params RemoveInstanceFromMonitorGroupDetailParam `json:"params"` // 详细参数
}

