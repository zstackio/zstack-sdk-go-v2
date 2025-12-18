// Copyright (c) ZStack.io, Inc.

package param

// CreateRootVolumeTemplateFromRootVolumeDetailParam CreateRootVolumeTemplateFromRootVolume详细参数
type CreateRootVolumeTemplateFromRootVolumeDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"guestOsType,omitempty"`
	rest []string `json:"backupStorageUuids,omitempty"`
	rest string `json:"rootVolumeUuid" validate:"required"` // 必填
	rest string `json:"platform,omitempty"`
	rest bool `json:"system,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest bool `json:"virtio,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateRootVolumeTemplateFromRootVolumeParam CreateRootVolumeTemplateFromRootVolume请求参数
type CreateRootVolumeTemplateFromRootVolumeParam struct {
	BaseParam
	Params CreateRootVolumeTemplateFromRootVolumeDetailParam `json:"params"` // 详细参数
}

