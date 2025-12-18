// Copyright (c) ZStack.io, Inc.

package param

// PrimaryStorageMigrateVolumeDetailParam PrimaryStorageMigrateVolume detail param
type PrimaryStorageMigrateVolumeDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	DstPrimaryStorageUuid string `json:"dstPrimaryStorageUuid" validate:"required"`
	VolumeProvisioningStrategy string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVolumeParam PrimaryStorageMigrateVolume request param
type PrimaryStorageMigrateVolumeParam struct {
	BaseParam
	Params PrimaryStorageMigrateVolumeDetailParam `json:"params"`
}
