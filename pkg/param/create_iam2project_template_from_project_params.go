// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectTemplateFromProjectDetailParam CreateIAM2ProjectTemplateFromProject detail param
type CreateIAM2ProjectTemplateFromProjectDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ProjectUuid string `json:"projectUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectTemplateFromProjectParam CreateIAM2ProjectTemplateFromProject request param
type CreateIAM2ProjectTemplateFromProjectParam struct {
	BaseParam
	Params CreateIAM2ProjectTemplateFromProjectDetailParam `json:"params"`
}
