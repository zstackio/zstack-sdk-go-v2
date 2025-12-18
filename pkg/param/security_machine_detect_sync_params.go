// Copyright (c) ZStack.io, Inc.

package param

// SecurityMachineDetectSyncDetailParam SecurityMachineDetectSync detail param
type SecurityMachineDetectSyncDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SecurityMachineDetectSyncParam SecurityMachineDetectSync request param
type SecurityMachineDetectSyncParam struct {
	BaseParam
	Params SecurityMachineDetectSyncDetailParam `json:"params"`
}
