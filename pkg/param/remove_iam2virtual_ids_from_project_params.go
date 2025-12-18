// Copyright (c) ZStack.io, Inc.

package param

// RemoveIAM2VirtualIDsFromProjectDetailParam RemoveIAM2VirtualIDsFromProject详细参数
type RemoveIAM2VirtualIDsFromProjectDetailParam struct {
	rest string `json:"projectUuid,omitempty"`
	rest []string `json:"virtualIDUuids" validate:"required"` // 必填
}

// RemoveIAM2VirtualIDsFromProjectParam RemoveIAM2VirtualIDsFromProject请求参数
type RemoveIAM2VirtualIDsFromProjectParam struct {
	BaseParam
	Params RemoveIAM2VirtualIDsFromProjectDetailParam `json:"params"` // 详细参数
}

