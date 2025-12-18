// Copyright (c) ZStack.io, Inc.

package param

// AddInstanceToMonitorGroupDetailParam AddInstanceToMonitorGroup详细参数
type AddInstanceToMonitorGroupDetailParam struct {
	rest string `json:"instanceUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddInstanceToMonitorGroupParam AddInstanceToMonitorGroup请求参数
type AddInstanceToMonitorGroupParam struct {
	BaseParam
	Params AddInstanceToMonitorGroupDetailParam `json:"params"` // 详细参数
}

