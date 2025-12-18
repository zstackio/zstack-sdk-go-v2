// Copyright (c) ZStack.io, Inc.

package param

// GetVRouterRouteTableDetailParam GetVRouterRouteTable detail param
type GetVRouterRouteTableDetailParam struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" validate:"required"`
}

// GetVRouterRouteTableParam GetVRouterRouteTable request param
type GetVRouterRouteTableParam struct {
	BaseParam
	Params GetVRouterRouteTableDetailParam `json:"params"`
}
