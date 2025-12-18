// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkRouterInterfaceIpDetailParam GetL3NetworkRouterInterfaceIp detail param
type GetL3NetworkRouterInterfaceIpDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
}

// GetL3NetworkRouterInterfaceIpParam GetL3NetworkRouterInterfaceIp request param
type GetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params GetL3NetworkRouterInterfaceIpDetailParam `json:"params"`
}
