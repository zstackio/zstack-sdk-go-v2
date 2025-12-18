// Copyright (c) ZStack.io, Inc.

package param

// StopVmInstanceDetailParam StopVmInstance detail param
type StopVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Type string `json:"type,omitempty"`
	StopHA string `json:"stopHA,omitempty"`
}

// StopVmInstanceParam StopVmInstance request param
type StopVmInstanceParam struct {
	BaseParam
	Params StopVmInstanceDetailParam `json:"params"`
}
