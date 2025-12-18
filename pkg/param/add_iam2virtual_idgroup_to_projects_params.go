// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDGroupToProjectsDetailParam AddIAM2VirtualIDGroupToProjects详细参数
type AddIAM2VirtualIDGroupToProjectsDetailParam struct {
	rest []interface{} `json:"structs,omitempty"`
}

// AddIAM2VirtualIDGroupToProjectsParam AddIAM2VirtualIDGroupToProjects请求参数
type AddIAM2VirtualIDGroupToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDGroupToProjectsDetailParam `json:"params"` // 详细参数
}

