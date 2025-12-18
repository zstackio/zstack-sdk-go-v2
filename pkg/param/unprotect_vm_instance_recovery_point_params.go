// Copyright (c) ZStack.io, Inc.

package param

// UnprotectVmInstanceRecoveryPointDetailParam UnprotectVmInstanceRecoveryPoint详细参数
type UnprotectVmInstanceRecoveryPointDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
}

// UnprotectVmInstanceRecoveryPointParam UnprotectVmInstanceRecoveryPoint请求参数
type UnprotectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnprotectVmInstanceRecoveryPointDetailParam `json:"params"` // 详细参数
}

