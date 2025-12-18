// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToOrganizationDetailParam AddIAM2VirtualIDsToOrganization detail param
type AddIAM2VirtualIDsToOrganizationDetailParam struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	OrganizationUuid string `json:"organizationUuid" validate:"required"`
}

// AddIAM2VirtualIDsToOrganizationParam AddIAM2VirtualIDsToOrganization request param
type AddIAM2VirtualIDsToOrganizationParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToOrganizationDetailParam `json:"params"`
}
