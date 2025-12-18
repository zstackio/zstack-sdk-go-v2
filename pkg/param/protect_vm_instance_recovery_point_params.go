// Copyright (c) ZStack.io, Inc.

package param

// ProtectVmInstanceRecoveryPointDetailParam ProtectVmInstanceRecoveryPoint详细参数
type ProtectVmInstanceRecoveryPointDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// ProtectVmInstanceRecoveryPointParam ProtectVmInstanceRecoveryPoint请求参数
type ProtectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params ProtectVmInstanceRecoveryPointDetailParam `json:"params"` // 详细参数
}

