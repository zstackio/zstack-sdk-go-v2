// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromProjectsDetailParam RemoveIAM2VirtualIDsFromProjects detail param
type RemoveIAM2VirtualIDsFromProjectsDetailParam struct {
	ProjectUuids []string `json:"projectUuids" validate:"required"`
	VirtualIDUuids []string `json:"virtualIDUuids" validate:"required"`
}

// RemoveIAM2VirtualIDsFromProjectsParam RemoveIAM2VirtualIDsFromProjects request param
type RemoveIAM2VirtualIDsFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectsDetailParam `json:"params"`
}
