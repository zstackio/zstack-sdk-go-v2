// Copyright (c) ZStack.io, Inc.

package param

// FailoverFaultToleranceVmDetailParam FailoverFaultToleranceVm detail param
type FailoverFaultToleranceVmDetailParam struct {
	FaultToleranceVmUuid string `json:"faultToleranceVmUuid" validate:"required"`
}

// FailoverFaultToleranceVmParam FailoverFaultToleranceVm request param
type FailoverFaultToleranceVmParam struct {
	BaseParam
	Params FailoverFaultToleranceVmDetailParam `json:"params"`
}
