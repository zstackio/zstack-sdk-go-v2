// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDGroupFromProjectsDetailParam RemoveIAM2VirtualIDGroupFromProjects详细参数
type RemoveIAM2VirtualIDGroupFromProjectsDetailParam struct {
	rest []string `json:"projectUuids,omitempty"`
	rest []string `json:"groupUuids,omitempty"`
}

// RemoveIAM2VirtualIDGroupFromProjectsParam RemoveIAM2VirtualIDGroupFromProjects请求参数
type RemoveIAM2VirtualIDGroupFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDGroupFromProjectsDetailParam `json:"params"` // 详细参数
}

