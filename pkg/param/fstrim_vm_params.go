// Copyright (c) ZStack.io, Inc.

package param

// FstrimVmDetailParam FstrimVm detail param
type FstrimVmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// FstrimVmParam FstrimVm request param
type FstrimVmParam struct {
	BaseParam
	Params FstrimVmDetailParam `json:"params"`
}
