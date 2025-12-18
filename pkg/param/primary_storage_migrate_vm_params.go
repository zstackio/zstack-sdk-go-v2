// Copyright (c) ZStack.io, Inc.

package param

// PrimaryStorageMigrateVmDetailParam PrimaryStorageMigrateVm detail param
type PrimaryStorageMigrateVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	DstHostUuid string `json:"dstHostUuid,omitempty"`
	WithDataVolumes bool `json:"withDataVolumes,omitempty"`
	DataVolumeUuids []string `json:"dataVolumeUuids,omitempty"`
	WithSnapshots bool `json:"withSnapshots,omitempty"`
	DownTime int `json:"downTime,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	Bandwidth int64 `json:"bandwidth,omitempty"`
	VolumeProvisioningStrategy string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVmParam PrimaryStorageMigrateVm request param
type PrimaryStorageMigrateVmParam struct {
	BaseParam
	Params PrimaryStorageMigrateVmDetailParam `json:"params"`
}
