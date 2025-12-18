// Copyright (c) ZStack.io, Inc.

package param

// RemoveAttributesFromIAM2OrganizationDetailParam RemoveAttributesFromIAM2Organization detail param
type RemoveAttributesFromIAM2OrganizationDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	AttributeUuids []string `json:"attributeUuids" validate:"required"`
}

// RemoveAttributesFromIAM2OrganizationParam RemoveAttributesFromIAM2Organization request param
type RemoveAttributesFromIAM2OrganizationParam struct {
	BaseParam
	Params RemoveAttributesFromIAM2OrganizationDetailParam `json:"params"`
}
