// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkIpStatisticDetailParam GetL3NetworkIpStatistic detail param
type GetL3NetworkIpStatisticDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	ResourceType string `json:"resourceType,omitempty"`
	Ip string `json:"ip,omitempty"`
	SortBy string `json:"sortBy,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
	Start int `json:"start,omitempty"`
	Limit int `json:"limit,omitempty"`
	ReplyWithCount bool `json:"replyWithCount,omitempty"`
}

// GetL3NetworkIpStatisticParam GetL3NetworkIpStatistic request param
type GetL3NetworkIpStatisticParam struct {
	BaseParam
	Params GetL3NetworkIpStatisticDetailParam `json:"params"`
}
