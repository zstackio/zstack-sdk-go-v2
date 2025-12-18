// Copyright (c) ZStack.io, Inc.

package param

// GetVmRDPDetailParam GetVmRDP detail param
type GetVmRDPDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmRDPParam GetVmRDP request param
type GetVmRDPParam struct {
	BaseParam
	Params GetVmRDPDetailParam `json:"params"`
}
