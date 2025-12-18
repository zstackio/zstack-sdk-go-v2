// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromOrganizationDetailParam RemoveIAM2VirtualIDsFromOrganization detail param
type RemoveIAM2VirtualIDsFromOrganizationDetailParam struct {
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	OrganizationUuid string `json:"organizationUuid" validate:"required"`
}

// RemoveIAM2VirtualIDsFromOrganizationParam RemoveIAM2VirtualIDsFromOrganization request param
type RemoveIAM2VirtualIDsFromOrganizationParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromOrganizationDetailParam `json:"params"`
}
