// Copyright (c) ZStack.io, Inc.

package param

// GetMonitorItemDetailParam GetMonitorItem详细参数
type GetMonitorItemDetailParam struct {
	rest string `json:"resourceType" validate:"required"` // 必填
}

// GetMonitorItemParam GetMonitorItem请求参数
type GetMonitorItemParam struct {
	BaseParam
	Params GetMonitorItemDetailParam `json:"params"` // 详细参数
}

