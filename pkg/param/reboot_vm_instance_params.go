// Copyright (c) ZStack.io, Inc.

package param

// RebootVmInstanceDetailParam RebootVmInstance detail param
type RebootVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RebootVmInstanceParam RebootVmInstance request param
type RebootVmInstanceParam struct {
	BaseParam
	Params RebootVmInstanceDetailParam `json:"params"`
}
