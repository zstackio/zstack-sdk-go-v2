// Copyright (c) ZStack.io, Inc.

package param

// CreateDataVolumeFromVolumeTemplateDetailParam CreateDataVolumeFromVolumeTemplate detail param
type CreateDataVolumeFromVolumeTemplateDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	HostUuid string `json:"hostUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDataVolumeFromVolumeTemplateParam CreateDataVolumeFromVolumeTemplate request param
type CreateDataVolumeFromVolumeTemplateParam struct {
	BaseParam
	Params CreateDataVolumeFromVolumeTemplateDetailParam `json:"params"`
}
