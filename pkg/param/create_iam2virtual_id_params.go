// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDDetailParam CreateIAM2VirtualID detail param
type CreateIAM2VirtualIDDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Description string `json:"description,omitempty"`
	Attributes []interface{} `json:"attributes,omitempty"`
	ProjectUuid string `json:"projectUuid,omitempty"`
	OrganizationUuid string `json:"organizationUuid,omitempty"`
	WithoutDefaultRole bool `json:"withoutDefaultRole,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDParam CreateIAM2VirtualID request param
type CreateIAM2VirtualIDParam struct {
	BaseParam
	Params CreateIAM2VirtualIDDetailParam `json:"params"`
}
