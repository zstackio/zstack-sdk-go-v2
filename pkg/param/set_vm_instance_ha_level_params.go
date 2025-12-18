// Copyright (c) ZStack.io, Inc.

package param

// SetVmInstanceHaLevelDetailParam SetVmInstanceHaLevel detail param
type SetVmInstanceHaLevelDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Level string `json:"level" validate:"required"`
}

// SetVmInstanceHaLevelParam SetVmInstanceHaLevel request param
type SetVmInstanceHaLevelParam struct {
	BaseParam
	Params SetVmInstanceHaLevelDetailParam `json:"params"`
}
