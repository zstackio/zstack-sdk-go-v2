// Copyright (c) ZStack.io, Inc.

package param

// UnmountVmInstanceRecoveryPointDetailParam UnmountVmInstanceRecoveryPoint详细参数
type UnmountVmInstanceRecoveryPointDetailParam struct {
	rest string `json:"vmUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
}

// UnmountVmInstanceRecoveryPointParam UnmountVmInstanceRecoveryPoint请求参数
type UnmountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params UnmountVmInstanceRecoveryPointDetailParam `json:"params"` // 详细参数
}

