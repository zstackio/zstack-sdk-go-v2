// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDGroupToProjectsDetailParam AddIAM2VirtualIDGroupToProjects detail param
type AddIAM2VirtualIDGroupToProjectsDetailParam struct {
	Structs []IAM2ProjectRoleRefStructParam `json:"structs,omitempty"`
}

// AddIAM2VirtualIDGroupToProjectsParam AddIAM2VirtualIDGroupToProjects request param
type AddIAM2VirtualIDGroupToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDGroupToProjectsDetailParam `json:"params"`
}
