// Copyright (c) ZStack.io, Inc.

package param

// FailoverFaultToleranceVmDetailParam FailoverFaultToleranceVm详细参数
type FailoverFaultToleranceVmDetailParam struct {
	rest string `json:"faultToleranceVmUuid" validate:"required"` // 必填
}

// FailoverFaultToleranceVmParam FailoverFaultToleranceVm请求参数
type FailoverFaultToleranceVmParam struct {
	BaseParam
	Params FailoverFaultToleranceVmDetailParam `json:"params"` // 详细参数
}

