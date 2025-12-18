// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeTemplateDetailParam CreateDataVolumeFromVolumeTemplate详细参数
type CreateDataVolumeFromVolumeTemplateDetailParam struct {
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeTemplateParam CreateDataVolumeFromVolumeTemplate请求参数
type CreateDataVolumeFromVolumeTemplateParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeTemplateDetailParam `json:"params"` // 详细参数
}

