// Copyright (c) ZStack.io, Inc.

package param

// SetVmRDPDetailParam SetVmRDP detail param
type SetVmRDPDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetVmRDPParam SetVmRDP request param
type SetVmRDPParam struct {
	BaseParam
	Params SetVmRDPDetailParam `json:"params"`
}
