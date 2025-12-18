// Copyright (c) ZStack.io, Inc.

package param

// SetServiceTypeOnHostNetworkInterfaceDetailParam SetServiceTypeOnHostNetworkInterface详细参数
type SetServiceTypeOnHostNetworkInterfaceDetailParam struct {
	rest []string `json:"interfaceUuids" validate:"required"` // 必填
	rest []int `json:"vlanIds,omitempty"`
	rest []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceParam SetServiceTypeOnHostNetworkInterface请求参数
type SetServiceTypeOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkInterfaceDetailParam `json:"params"` // 详细参数
}

