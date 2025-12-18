// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToProjectsDetailParam AddIAM2VirtualIDsToProjects详细参数
type AddIAM2VirtualIDsToProjectsDetailParam struct {
	rest []string `json:"projectUuids" validate:"required"` // 必填
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectsParam AddIAM2VirtualIDsToProjects请求参数
type AddIAM2VirtualIDsToProjectsParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectsDetailParam `json:"params"` // 详细参数
}

