// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2OrganizationAttributeDetailParam UpdateIAM2OrganizationAttribute detail param
type UpdateIAM2OrganizationAttributeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2OrganizationAttributeParam UpdateIAM2OrganizationAttribute request param
type UpdateIAM2OrganizationAttributeParam struct {
	BaseParam
	Params UpdateIAM2OrganizationAttributeDetailParam `json:"params"`
}
