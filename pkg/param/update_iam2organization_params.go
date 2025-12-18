// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2OrganizationDetailParam UpdateIAM2Organization detail param
type UpdateIAM2OrganizationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateIAM2OrganizationParam UpdateIAM2Organization request param
type UpdateIAM2OrganizationParam struct {
	BaseParam
	Params UpdateIAM2OrganizationDetailParam `json:"params"`
}
