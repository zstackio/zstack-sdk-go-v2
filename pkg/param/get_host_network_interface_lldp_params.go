// Copyright (c) ZStack.io, Inc.

package param

// GetHostNetworkInterfaceLldpDetailParam GetHostNetworkInterfaceLldp detail param
type GetHostNetworkInterfaceLldpDetailParam struct {
	InterfaceUuid string `json:"interfaceUuid" validate:"required"`
}

// GetHostNetworkInterfaceLldpParam GetHostNetworkInterfaceLldp request param
type GetHostNetworkInterfaceLldpParam struct {
	BaseParam
	Params GetHostNetworkInterfaceLldpDetailParam `json:"params"`
}
