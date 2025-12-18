// Copyright (c) ZStack.io, Inc.

package param

// MountVmInstanceRecoveryPointDetailParam MountVmInstanceRecoveryPoint detail param
type MountVmInstanceRecoveryPointDetailParam struct {
	VmUuid string `json:"vmUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	Https bool `json:"https,omitempty"`
}

// MountVmInstanceRecoveryPointParam MountVmInstanceRecoveryPoint request param
type MountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params MountVmInstanceRecoveryPointDetailParam `json:"params"`
}
