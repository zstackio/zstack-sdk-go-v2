// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkIpStatisticDetailParam GetL3NetworkIpStatistic详细参数
type GetL3NetworkIpStatisticDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"resourceType,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest int `json:"start,omitempty"`
	rest int `json:"limit,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
}

// GetL3NetworkIpStatisticParam GetL3NetworkIpStatistic请求参数
type GetL3NetworkIpStatisticParam struct {
	BaseParam
	Params GetL3NetworkIpStatisticDetailParam `json:"params"` // 详细参数
}

