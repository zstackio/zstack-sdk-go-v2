// Copyright (c) ZStack.io, Inc.

package param

// AddIAM2VirtualIDsToGroupDetailParam AddIAM2VirtualIDsToGroup详细参数
type AddIAM2VirtualIDsToGroupDetailParam struct {
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// AddIAM2VirtualIDsToGroupParam AddIAM2VirtualIDsToGroup请求参数
type AddIAM2VirtualIDsToGroupParam struct {
	BaseParam
	Params AddIAM2VirtualIDsToGroupDetailParam `json:"params"` // 详细参数
}

