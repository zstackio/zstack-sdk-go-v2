// Copyright (c) ZStack.io, Inc.

package param

// ChangeIAM2OrganizationParentDetailParam ChangeIAM2OrganizationParent detail param
type ChangeIAM2OrganizationParentDetailParam struct {
	ParentUuid string `json:"parentUuid" validate:"required"`
	ChildrenUuids []string `json:"childrenUuids" validate:"required"`
}

// ChangeIAM2OrganizationParentParam ChangeIAM2OrganizationParent request param
type ChangeIAM2OrganizationParentParam struct {
	BaseParam
	Params ChangeIAM2OrganizationParentDetailParam `json:"params"`
}
