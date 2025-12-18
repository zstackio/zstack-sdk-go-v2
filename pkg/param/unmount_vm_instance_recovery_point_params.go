// Copyright (c) ZStack.io, Inc.

package param

// UnmountVmInstanceRecoveryPointDetailParam UnmountVmInstanceRecoveryPoint detail param
type UnmountVmInstanceRecoveryPointDetailParam struct {
	VmUuid string `json:"vmUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
}

// UnmountVmInstanceRecoveryPointParam UnmountVmInstanceRecoveryPoint request param
type UnmountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnmountVmInstanceRecoveryPointDetailParam `json:"params"`
}
