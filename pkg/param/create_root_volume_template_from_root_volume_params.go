// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromRootVolumeDetailParam CreateRootVolumeTemplateFromRootVolume detail param
type CreateRootVolumeTemplateFromRootVolumeDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	RootVolumeUuid string `json:"rootVolumeUuid" validate:"required"`
	Platform string `json:"platform,omitempty"`
	System bool `json:"system,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromRootVolumeParam CreateRootVolumeTemplateFromRootVolume request param
type CreateRootVolumeTemplateFromRootVolumeParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromRootVolumeDetailParam `json:"params"`
}
