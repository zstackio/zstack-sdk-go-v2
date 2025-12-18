// Copyright (c) ZStack.io, Inc.

package param

// GetInterfaceServiceTypeStatisticDetailParam GetInterfaceServiceTypeStatistic详细参数
type GetInterfaceServiceTypeStatisticDetailParam struct {
	rest string `json:"interfaceUuid,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"interfaceType,omitempty"`
	rest []string `json:"serviceType,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest int `json:"start,omitempty"`
	rest int `json:"limit,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
}

// GetInterfaceServiceTypeStatisticParam GetInterfaceServiceTypeStatistic请求参数
type GetInterfaceServiceTypeStatisticParam struct {
	BaseParam
	Params GetInterfaceServiceTypeStatisticDetailParam `json:"params"` // 详细参数
}

