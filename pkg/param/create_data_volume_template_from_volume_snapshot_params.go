// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeTemplateFromVolumeSnapshotDetailParam CreateDataVolumeTemplateFromVolumeSnapshot detail param
type CreateDataVolumeTemplateFromVolumeSnapshotDetailParam struct {
	SnapshotUuid string `json:"snapshotUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeSnapshotParam CreateDataVolumeTemplateFromVolumeSnapshot request param
type CreateDataVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeSnapshotDetailParam `json:"params"`
}
