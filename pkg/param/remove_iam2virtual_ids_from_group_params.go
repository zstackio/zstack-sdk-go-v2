// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromGroupDetailParam RemoveIAM2VirtualIDsFromGroup详细参数
type RemoveIAM2VirtualIDsFromGroupDetailParam struct {
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// RemoveIAM2VirtualIDsFromGroupParam RemoveIAM2VirtualIDsFromGroup请求参数
type RemoveIAM2VirtualIDsFromGroupParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromGroupDetailParam `json:"params"` // 详细参数
}

