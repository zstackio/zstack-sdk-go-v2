// Copyright (c) ZStack.io, Inc.

package param

// SetL3NetworkRouterInterfaceIpDetailParam SetL3NetworkRouterInterfaceIp detail param
type SetL3NetworkRouterInterfaceIpDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	RouterInterfaceIp string `json:"routerInterfaceIp" validate:"required"`
}

// SetL3NetworkRouterInterfaceIpParam SetL3NetworkRouterInterfaceIp request param
type SetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params SetL3NetworkRouterInterfaceIpDetailParam `json:"params"`
}
