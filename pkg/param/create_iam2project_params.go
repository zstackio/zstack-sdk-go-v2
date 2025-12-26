// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectDetailParam CreateIAM2Project detail param
type CreateIAM2ProjectDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	RoleUuids []string `json:"roleUuids,omitempty"`
	ResourceTemplates []string `json:"resourceTemplates,omitempty"`
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	LinkAccountUuid string `json:"linkAccountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectParam CreateIAM2Project request param
type CreateIAM2ProjectParam struct {
	BaseParam
	Params CreateIAM2ProjectDetailParam `json:"params"`
}
