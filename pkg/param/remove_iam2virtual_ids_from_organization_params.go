// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromOrganizationDetailParam RemoveIAM2VirtualIDsFromOrganization详细参数
type RemoveIAM2VirtualIDsFromOrganizationDetailParam struct {
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest string `json:"organizationUuid" validate:"required"` // 必填
}

// RemoveIAM2VirtualIDsFromOrganizationParam RemoveIAM2VirtualIDsFromOrganization请求参数
type RemoveIAM2VirtualIDsFromOrganizationParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromOrganizationDetailParam `json:"params"` // 详细参数
}

