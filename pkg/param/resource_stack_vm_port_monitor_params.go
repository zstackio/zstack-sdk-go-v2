// Copyright (c) ZStack.io, Inc.

package param

// DeleteResourceStackVmPortMonitorDetailParam DeleteResourceStackVmPortMonitor详细参数
type DeleteResourceStackVmPortMonitorDetailParam struct {
	rest string `json:"stackUuid,omitempty"`
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int `json:"port,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteResourceStackVmPortMonitorParam DeleteResourceStackVmPortMonitor请求参数
type DeleteResourceStackVmPortMonitorParam struct {
	BaseParam
	Params DeleteResourceStackVmPortMonitorDetailParam `json:"params"` // 详细参数
}

