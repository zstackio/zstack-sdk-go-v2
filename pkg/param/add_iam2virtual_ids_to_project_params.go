// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToProjectDetailParam AddIAM2VirtualIDsToProject detail param
type AddIAM2VirtualIDsToProjectDetailParam struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectParam AddIAM2VirtualIDsToProject request param
type AddIAM2VirtualIDsToProjectParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectDetailParam `json:"params"`
}
