// Copyright (c) ZStack.io, Inc.

package param

// SetServiceTypeOnHostNetworkInterfaceDetailParam SetServiceTypeOnHostNetworkInterface detail param
type SetServiceTypeOnHostNetworkInterfaceDetailParam struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	VlanIds []int `json:"vlanIds,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceParam SetServiceTypeOnHostNetworkInterface request param
type SetServiceTypeOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetServiceTypeOnHostNetworkInterfaceDetailParam `json:"params"`
}
