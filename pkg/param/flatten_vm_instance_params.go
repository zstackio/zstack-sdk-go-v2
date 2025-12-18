// Copyright (c) ZStack.io, Inc.

package param

// FlattenVmInstanceDetailParam FlattenVmInstance detail param
type FlattenVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Full bool `json:"full,omitempty"`
	DryRun bool `json:"dryRun,omitempty"`
}

// FlattenVmInstanceParam FlattenVmInstance request param
type FlattenVmInstanceParam struct {
	BaseParam
	Params FlattenVmInstanceDetailParam `json:"params"`
}
