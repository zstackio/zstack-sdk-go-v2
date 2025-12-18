// Copyright (c) ZStack.io, Inc.

package param

// DetachVRouterRouteTableFromVRouterDetailParam DetachVRouterRouteTableFromVRouter详细参数
type DetachVRouterRouteTableFromVRouterDetailParam struct {
	rest string `json:"routeTableUuid" validate:"required"` // 必填
	rest string `json:"virtualRouterVmUuid" validate:"required"` // 必填
}

// DetachVRouterRouteTableFromVRouterParam DetachVRouterRouteTableFromVRouter请求参数
type DetachVRouterRouteTableFromVRouterParam struct {
	BaseParam
	Params DetachVRouterRouteTableFromVRouterDetailParam `json:"params"` // 详细参数
}

