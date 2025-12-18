// Copyright (c) ZStack.io, Inc.

package param

// GetFaultToleranceVmsDetailParam GetFaultToleranceVms detail param
type GetFaultToleranceVmsDetailParam struct {
	FaultToleranceVmUuid string `json:"faultToleranceVmUuid" validate:"required"`
}

// GetFaultToleranceVmsParam GetFaultToleranceVms request param
type GetFaultToleranceVmsParam struct {
	BaseParam
	Params GetFaultToleranceVmsDetailParam `json:"params"`
}
