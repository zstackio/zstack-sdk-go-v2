// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDGroupFromProjectsDetailParam RemoveIAM2VirtualIDGroupFromProjects detail param
type RemoveIAM2VirtualIDGroupFromProjectsDetailParam struct {
	ProjectUuids []string `json:"projectUuids,omitempty"`
	GroupUuids []string `json:"groupUuids,omitempty"`
}

// RemoveIAM2VirtualIDGroupFromProjectsParam RemoveIAM2VirtualIDGroupFromProjects request param
type RemoveIAM2VirtualIDGroupFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDGroupFromProjectsDetailParam `json:"params"`
}
