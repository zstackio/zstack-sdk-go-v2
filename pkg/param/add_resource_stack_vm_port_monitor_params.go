// Copyright (c) ZStack.io, Inc.

package param

// AddResourceStackVmPortMonitorDetailParam AddResourceStackVmPortMonitor详细参数
type AddResourceStackVmPortMonitorDetailParam struct {
	rest string `json:"stackUuid,omitempty"`
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int `json:"port" validate:"required"` // 必填
}

// AddResourceStackVmPortMonitorParam AddResourceStackVmPortMonitor请求参数
type AddResourceStackVmPortMonitorParam struct {
	BaseParam
	Params AddResourceStackVmPortMonitorDetailParam `json:"params"` // 详细参数
}

