// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromVolumeSnapshotDetailParam CreateRootVolumeTemplateFromVolumeSnapshot detail param
type CreateRootVolumeTemplateFromVolumeSnapshotDetailParam struct {
	SnapshotUuid string `json:"snapshotUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromVolumeSnapshotParam CreateRootVolumeTemplateFromVolumeSnapshot request param
type CreateRootVolumeTemplateFromVolumeSnapshotParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromVolumeSnapshotDetailParam `json:"params"`
}
