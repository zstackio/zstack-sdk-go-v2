// Copyright (c) ZStack.io, Inc.

package param

// AddImageDetailParam AddImage detail param
type AddImageDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	MediaType string `json:"mediaType,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	System bool `json:"system,omitempty"`
	Format string `json:"format,omitempty"`
	Platform string `json:"platform,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids" validate:"required"`
	Type string `json:"type,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddImageParam AddImage request param
type AddImageParam struct {
	BaseParam
	Params AddImageDetailParam `json:"params"`
}
