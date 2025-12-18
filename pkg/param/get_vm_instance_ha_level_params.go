// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceHaLevelDetailParam GetVmInstanceHaLevel detail param
type GetVmInstanceHaLevelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmInstanceHaLevelParam GetVmInstanceHaLevel request param
type GetVmInstanceHaLevelParam struct {
	BaseParam
	Params GetVmInstanceHaLevelDetailParam `json:"params"`
}
