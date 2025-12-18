// Copyright (c) ZStack.io, Inc.

package param

// GetInterfaceServiceTypeStatisticDetailParam GetInterfaceServiceTypeStatistic detail param
type GetInterfaceServiceTypeStatisticDetailParam struct {
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	InterfaceType string `json:"interfaceType,omitempty"`
	ServiceType []string `json:"serviceType,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	SortBy string `json:"sortBy,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
	Start int `json:"start,omitempty"`
	Limit int `json:"limit,omitempty"`
	ReplyWithCount bool `json:"replyWithCount,omitempty"`
}

// GetInterfaceServiceTypeStatisticParam GetInterfaceServiceTypeStatistic request param
type GetInterfaceServiceTypeStatisticParam struct {
	BaseParam
	Params GetInterfaceServiceTypeStatisticDetailParam `json:"params"`
}
