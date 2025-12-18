// Copyright (c) ZStack.io, Inc.

package param

// PauseVmInstanceDetailParam PauseVmInstance detail param
type PauseVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// PauseVmInstanceParam PauseVmInstance request param
type PauseVmInstanceParam struct {
	BaseParam
	Params PauseVmInstanceDetailParam `json:"params"`
}
