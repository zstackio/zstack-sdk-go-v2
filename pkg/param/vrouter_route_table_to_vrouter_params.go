// Copyright (c) ZStack.io, Inc.

package param

// AttachVRouterRouteTableToVRouterDetailParam AttachVRouterRouteTableToVRouter详细参数
type AttachVRouterRouteTableToVRouterDetailParam struct {
	rest string `json:"routeTableUuid" validate:"required"` // 必填
	rest string `json:"virtualRouterVmUuid" validate:"required"` // 必填
}

// AttachVRouterRouteTableToVRouterParam AttachVRouterRouteTableToVRouter请求参数
type AttachVRouterRouteTableToVRouterParam struct {
	BaseParam
	Params AttachVRouterRouteTableToVRouterDetailParam `json:"params"` // 详细参数
}

