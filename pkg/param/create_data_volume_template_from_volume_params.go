// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeTemplateFromVolumeDetailParam CreateDataVolumeTemplateFromVolume detail param
type CreateDataVolumeTemplateFromVolumeDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeParam CreateDataVolumeTemplateFromVolume request param
type CreateDataVolumeTemplateFromVolumeParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeDetailParam `json:"params"`
}
