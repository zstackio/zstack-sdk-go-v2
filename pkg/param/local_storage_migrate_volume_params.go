// Copyright (c) ZStack.io, Inc.

package param

// LocalStorageMigrateVolumeDetailParam LocalStorageMigrateVolume详细参数
type LocalStorageMigrateVolumeDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"destHostUuid" validate:"required"` // 必填
}

// LocalStorageMigrateVolumeParam LocalStorageMigrateVolume请求参数
type LocalStorageMigrateVolumeParam struct {
	BaseParam
	Params LocalStorageMigrateVolumeDetailParam `json:"params"` // 详细参数
}

