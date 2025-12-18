// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostNetworkInterfaceDetailParam UpdateHostNetworkInterface detail param
type UpdateHostNetworkInterfaceDetailParam struct {
	InterfaceUuid string `json:"interfaceUuid" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdateHostNetworkInterfaceParam UpdateHostNetworkInterface request param
type UpdateHostNetworkInterfaceParam struct {
	BaseParam
	Params UpdateHostNetworkInterfaceDetailParam `json:"params"`
}
