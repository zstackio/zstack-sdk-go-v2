// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToProjectDetailParam AddIAM2VirtualIDsToProject详细参数
type AddIAM2VirtualIDsToProjectDetailParam struct {
	rest string `json:"projectUuid,omitempty"`
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest []string `json:"roleUuids,omitempty"`
}

// AddIAM2VirtualIDsToProjectParam AddIAM2VirtualIDsToProject请求参数
type AddIAM2VirtualIDsToProjectParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToProjectDetailParam `json:"params"` // 详细参数
}

