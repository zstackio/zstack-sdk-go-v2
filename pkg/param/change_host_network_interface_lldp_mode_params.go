// Copyright (c) ZStack.io, Inc.

package param

// ChangeHostNetworkInterfaceLldpModeDetailParam ChangeHostNetworkInterfaceLldpMode detail param
type ChangeHostNetworkInterfaceLldpModeDetailParam struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	Mode string `json:"mode,omitempty"`
}

// ChangeHostNetworkInterfaceLldpModeParam ChangeHostNetworkInterfaceLldpMode request param
type ChangeHostNetworkInterfaceLldpModeParam struct {
	BaseParam
	Params ChangeHostNetworkInterfaceLldpModeDetailParam `json:"params"`
}
