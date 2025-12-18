// Copyright (c) ZStack.io, Inc.

package param

// SetVmMonitorNumberDetailParam SetVmMonitorNumber详细参数
type SetVmMonitorNumberDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"monitorNumber" validate:"required"` // 必填
}

// SetVmMonitorNumberParam SetVmMonitorNumber请求参数
type SetVmMonitorNumberParam struct {
	BaseParam
	Params SetVmMonitorNumberDetailParam `json:"params"` // 详细参数
}

