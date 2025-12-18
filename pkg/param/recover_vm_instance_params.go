// Copyright (c) ZStack.io, Inc.

package param

// RecoverVmInstanceDetailParam RecoverVmInstance detail param
type RecoverVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoverVmInstanceParam RecoverVmInstance request param
type RecoverVmInstanceParam struct {
	BaseParam
	Params RecoverVmInstanceDetailParam `json:"params"`
}
