// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2VirtualIDInGroupDetailParam GetIAM2VirtualIDInGroup详细参数
type GetIAM2VirtualIDInGroupDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
}

// GetIAM2VirtualIDInGroupParam GetIAM2VirtualIDInGroup请求参数
type GetIAM2VirtualIDInGroupParam struct {
	BaseParam
	Params GetIAM2VirtualIDInGroupDetailParam `json:"params"` // 详细参数
}

