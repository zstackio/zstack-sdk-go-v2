// Copyright (c) ZStack.io, Inc.

package param

// GetVmMonitorNumberDetailParam GetVmMonitorNumber详细参数
type GetVmMonitorNumberDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmMonitorNumberParam GetVmMonitorNumber请求参数
type GetVmMonitorNumberParam struct {
	BaseParam
	Params GetVmMonitorNumberDetailParam `json:"params"` // 详细参数
}

