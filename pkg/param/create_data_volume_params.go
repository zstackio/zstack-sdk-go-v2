// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeDetailParam CreateDataVolume detail param
type CreateDataVolumeDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	DiskOfferingUuid string `json:"diskOfferingUuid,omitempty"`
	DiskSize int64 `json:"diskSize,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeParam CreateDataVolume request param
type CreateDataVolumeParam struct {
	BaseParam
	Params CreateDataVolumeDetailParam `json:"params"`
}
