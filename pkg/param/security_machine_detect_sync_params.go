// Copyright (c) ZStack.io, Inc.

package param

// SecurityMachineDetectSyncDetailParam SecurityMachineDetectSync详细参数
type SecurityMachineDetectSyncDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// SecurityMachineDetectSyncParam SecurityMachineDetectSync请求参数
type SecurityMachineDetectSyncParam struct {
	BaseParam
	Params SecurityMachineDetectSyncDetailParam `json:"params"` // 详细参数
}

