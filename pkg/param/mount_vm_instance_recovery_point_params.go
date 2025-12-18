// Copyright (c) ZStack.io, Inc.

package param

// MountVmInstanceRecoveryPointDetailParam MountVmInstanceRecoveryPoint详细参数
type MountVmInstanceRecoveryPointDetailParam struct {
	rest string `json:"vmUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
	rest bool `json:"https,omitempty"`
}

// MountVmInstanceRecoveryPointParam MountVmInstanceRecoveryPoint请求参数
type MountVmInstanceRecoveryPointParam struct {
	BaseParam
	Params MountVmInstanceRecoveryPointDetailParam `json:"params"` // 详细参数
}

