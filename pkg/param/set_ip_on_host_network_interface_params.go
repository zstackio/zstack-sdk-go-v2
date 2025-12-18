// Copyright (c) ZStack.io, Inc.

package param

// SetIpOnHostNetworkInterfaceDetailParam SetIpOnHostNetworkInterface detail param
type SetIpOnHostNetworkInterfaceDetailParam struct {
	InterfaceUuid string `json:"interfaceUuid" validate:"required"`
	IpAddress string `json:"ipAddress,omitempty"`
	Netmask string `json:"netmask,omitempty"`
}

// SetIpOnHostNetworkInterfaceParam SetIpOnHostNetworkInterface request param
type SetIpOnHostNetworkInterfaceParam struct {
	BaseParam
	Params SetIpOnHostNetworkInterfaceDetailParam `json:"params"`
}
