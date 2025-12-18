// Copyright (c) ZStack.io, Inc.

package param

// ProtectVmInstanceRecoveryPointDetailParam ProtectVmInstanceRecoveryPoint detail param
type ProtectVmInstanceRecoveryPointDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	Description string `json:"description,omitempty"`
}

// ProtectVmInstanceRecoveryPointParam ProtectVmInstanceRecoveryPoint request param
type ProtectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params ProtectVmInstanceRecoveryPointDetailParam `json:"params"`
}
