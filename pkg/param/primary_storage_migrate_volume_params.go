// Copyright (c) ZStack.io, Inc.

package param

// PrimaryStorageMigrateVolumeDetailParam PrimaryStorageMigrateVolume详细参数
type PrimaryStorageMigrateVolumeDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"dstPrimaryStorageUuid" validate:"required"` // 必填
	rest string `json:"volumeProvisioningStrategy,omitempty"`
}

// PrimaryStorageMigrateVolumeParam PrimaryStorageMigrateVolume请求参数
type PrimaryStorageMigrateVolumeParam struct {
	BaseParam
	Params PrimaryStorageMigrateVolumeDetailParam `json:"params"` // 详细参数
}

