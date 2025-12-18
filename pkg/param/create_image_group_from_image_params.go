// Copyright (c) ZStack.io, Inc.

package param

// CreateImageGroupFromImageDetailParam CreateImageGroupFromImage detail param
type CreateImageGroupFromImageDetailParam struct {
	Name string `json:"name" validate:"required"`
	RootVolumeTemplateUuid string `json:"rootVolumeTemplateUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	DataVolumeTemplateUuids []string `json:"dataVolumeTemplateUuids,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageGroupFromImageParam CreateImageGroupFromImage request param
type CreateImageGroupFromImageParam struct {
	BaseParam
	Params CreateImageGroupFromImageDetailParam `json:"params"`
}
