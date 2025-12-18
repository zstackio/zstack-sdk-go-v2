// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectFromTemplateDetailParam CreateIAM2ProjectFromTemplate detail param
type CreateIAM2ProjectFromTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	TemplateUuid string `json:"templateUuid" validate:"required"`
	RoleUuids []string `json:"roleUuids,omitempty"`
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	ResourceTemplates []string `json:"resourceTemplates,omitempty"`
	LinkAccountUuid string `json:"linkAccountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectFromTemplateParam CreateIAM2ProjectFromTemplate request param
type CreateIAM2ProjectFromTemplateParam struct {
	BaseParam
	Params CreateIAM2ProjectFromTemplateDetailParam `json:"params"`
}
