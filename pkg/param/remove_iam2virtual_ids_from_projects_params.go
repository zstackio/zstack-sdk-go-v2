// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromProjectsDetailParam RemoveIAM2VirtualIDsFromProjects详细参数
type RemoveIAM2VirtualIDsFromProjectsDetailParam struct {
	rest []string `json:"projectUuids" validate:"required"` // 必填
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
}

// RemoveIAM2VirtualIDsFromProjectsParam RemoveIAM2VirtualIDsFromProjects请求参数
type RemoveIAM2VirtualIDsFromProjectsParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectsDetailParam `json:"params"` // 详细参数
}

