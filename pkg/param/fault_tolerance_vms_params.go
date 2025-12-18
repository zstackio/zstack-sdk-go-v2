// Copyright (c) ZStack.io, Inc.

package param

// GetFaultToleranceVmsDetailParam GetFaultToleranceVms详细参数
type GetFaultToleranceVmsDetailParam struct {
	rest string `json:"faultToleranceVmUuid" validate:"required"` // 必填
}

// GetFaultToleranceVmsParam GetFaultToleranceVms请求参数
type GetFaultToleranceVmsParam struct {
	BaseParam
	Params GetFaultToleranceVmsDetailParam `json:"params"` // 详细参数
}

