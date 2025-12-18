// Copyright (c) ZStack.io, Inc.

package param

// AttachVRouterRouteTableToVRouterDetailParam AttachVRouterRouteTableToVRouter detail param
type AttachVRouterRouteTableToVRouterDetailParam struct {
	RouteTableUuid string `json:"routeTableUuid" validate:"required"`
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" validate:"required"`
}

// AttachVRouterRouteTableToVRouterParam AttachVRouterRouteTableToVRouter request param
type AttachVRouterRouteTableToVRouterParam struct {
	BaseParam
	Params AttachVRouterRouteTableToVRouterDetailParam `json:"params"`
}
