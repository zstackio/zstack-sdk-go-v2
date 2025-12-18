// Copyright (c) ZStack.io, Inc.

package param

// LocalStorageMigrateVolumeDetailParam LocalStorageMigrateVolume detail param
type LocalStorageMigrateVolumeDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	DestHostUuid string `json:"destHostUuid" validate:"required"`
}

// LocalStorageMigrateVolumeParam LocalStorageMigrateVolume request param
type LocalStorageMigrateVolumeParam struct {
	BaseParam
	Params LocalStorageMigrateVolumeDetailParam `json:"params"`
}
