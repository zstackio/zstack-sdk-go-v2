// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2OrganizationDetailParam CreateIAM2Organization detail param
type CreateIAM2OrganizationDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	SrcType string `json:"srcType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2OrganizationParam CreateIAM2Organization request param
type CreateIAM2OrganizationParam struct {
	BaseParam
	Params CreateIAM2OrganizationDetailParam `json:"params"`
}
