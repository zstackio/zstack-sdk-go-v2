// Copyright (c) ZStack.io, Inc.

package param

// RevertVmFromCdpBackupDetailParam RevertVmFromCdpBackup detail param
type RevertVmFromCdpBackupDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	GroupId int64 `json:"groupId" validate:"required"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume string `json:"primaryStorageUuidForDataVolume,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	UseExistingVolume bool `json:"useExistingVolume,omitempty"`
	RecoverBandwidth int64 `json:"recoverBandwidth,omitempty"`
}

// RevertVmFromCdpBackupParam RevertVmFromCdpBackup request param
type RevertVmFromCdpBackupParam struct {
	BaseParam
	Params RevertVmFromCdpBackupDetailParam `json:"params"`
}
