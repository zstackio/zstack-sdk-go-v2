// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromCdpBackupDetailParam RevertVmFromCdpBackup详细参数
type RevertVmFromCdpBackupDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest int64 `json:"groupId" validate:"required"` // 必填
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"primaryStorageUuidForDataVolume,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest bool `json:"useExistingVolume,omitempty"`
	rest int64 `json:"recoverBandwidth,omitempty"`
}

// RevertVmFromCdpBackupParam RevertVmFromCdpBackup请求参数
type RevertVmFromCdpBackupParam struct {
	BaseParam
	Params RevertVmFromCdpBackupDetailParam `json:"params"` // 详细参数
}

