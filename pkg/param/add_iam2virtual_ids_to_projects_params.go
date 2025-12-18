// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToProjectsDetailParam AddIAM2VirtualIDsToProjects detail param
type AddIAM2VirtualIDsToProjectsDetailParam struct {
	ProjectUuids []string `json:"projectUuids" validate:"required"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectsParam AddIAM2VirtualIDsToProjects request param
type AddIAM2VirtualIDsToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectsDetailParam `json:"params"`
}
