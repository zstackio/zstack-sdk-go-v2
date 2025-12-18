// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToOrganizationDetailParam AddIAM2VirtualIDsToOrganization详细参数
type AddIAM2VirtualIDsToOrganizationDetailParam struct {
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest string `json:"organizationUuid" validate:"required"` // 必填
}

// AddIAM2VirtualIDsToOrganizationParam AddIAM2VirtualIDsToOrganization请求参数
type AddIAM2VirtualIDsToOrganizationParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToOrganizationDetailParam `json:"params"` // 详细参数
}

