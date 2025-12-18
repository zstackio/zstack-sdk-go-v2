// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmInstanceHaLevelDetailParam DeleteVmInstanceHaLevel detail param
type DeleteVmInstanceHaLevelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteVmInstanceHaLevelParam DeleteVmInstanceHaLevel request param
type DeleteVmInstanceHaLevelParam struct {
	BaseParam
	Params DeleteVmInstanceHaLevelDetailParam `json:"params"`
}
