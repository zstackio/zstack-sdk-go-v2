// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeTemplateFromVolumeDetailParam CreateDataVolumeTemplateFromVolume详细参数
type CreateDataVolumeTemplateFromVolumeDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest []string `json:"backupStorageUuids,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeTemplateFromVolumeParam CreateDataVolumeTemplateFromVolume请求参数
type CreateDataVolumeTemplateFromVolumeParam struct {
	BaseParam
	Params CreateDataVolumeTemplateFromVolumeDetailParam `json:"params"` // 详细参数
}

