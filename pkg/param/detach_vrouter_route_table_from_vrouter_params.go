// Copyright (c) ZStack.io, Inc.

package param

// DetachVRouterRouteTableFromVRouterDetailParam DetachVRouterRouteTableFromVRouter detail param
type DetachVRouterRouteTableFromVRouterDetailParam struct {
	RouteTableUuid string `json:"routeTableUuid" validate:"required"`
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" validate:"required"`
}

// DetachVRouterRouteTableFromVRouterParam DetachVRouterRouteTableFromVRouter request param
type DetachVRouterRouteTableFromVRouterParam struct {
	BaseParam
	Params DetachVRouterRouteTableFromVRouterDetailParam `json:"params"`
}
