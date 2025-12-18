// Copyright (c) ZStack.io, Inc.

package param

// UnprotectVmInstanceRecoveryPointDetailParam UnprotectVmInstanceRecoveryPoint detail param
type UnprotectVmInstanceRecoveryPointDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
}

// UnprotectVmInstanceRecoveryPointParam UnprotectVmInstanceRecoveryPoint request param
type UnprotectVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnprotectVmInstanceRecoveryPointDetailParam `json:"params"`
}
