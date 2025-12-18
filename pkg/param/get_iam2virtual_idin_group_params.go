// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2VirtualIDInGroupDetailParam GetIAM2VirtualIDInGroup detail param
type GetIAM2VirtualIDInGroupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
	Count bool `json:"count,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
	SortBy string `json:"sortBy,omitempty"`
	ReplyWithCount bool `json:"replyWithCount,omitempty"`
}

// GetIAM2VirtualIDInGroupParam GetIAM2VirtualIDInGroup request param
type GetIAM2VirtualIDInGroupParam struct {
	BaseParam
	Params GetIAM2VirtualIDInGroupDetailParam `json:"params"`
}
