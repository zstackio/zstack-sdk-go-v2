// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ContainerImageInventoryView ContainerImage
type ContainerImageInventoryView struct {
	EndpointUuid string `json:"endpointUuid,omitempty"`
	ImageTag string `json:"imageTag,omitempty"`
	RegistryUrl string `json:"registryUrl,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Url string `json:"url,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Type string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Format string `json:"format,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	BackupStorageRefs []ImageBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	SystemTags []SystemTagInventoryView `json:"systemTags,omitempty"`
}

// QueryContainerImageView QueryContainerImage
type QueryContainerImageView struct {
	Inventories []ContainerImageInventoryView `json:"inventories,omitempty"`
}

