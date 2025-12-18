// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2OrganizationParentDetailParam ChangeIAM2OrganizationParent详细参数
type ChangeIAM2OrganizationParentDetailParam struct {
	rest string `json:"parentUuid" validate:"required"` // 必填
	rest []string `json:"childrenUuids" validate:"required"` // 必填
}

// ChangeIAM2OrganizationParentParam ChangeIAM2OrganizationParent请求参数
type ChangeIAM2OrganizationParentParam struct {
	BaseParam
	Params ChangeIAM2OrganizationParentDetailParam `json:"params"` // 详细参数
}

