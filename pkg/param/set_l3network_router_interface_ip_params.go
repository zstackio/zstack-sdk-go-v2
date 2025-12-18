// Copyright (c) ZStack.io, Inc.

package param

// SetL3NetworkRouterInterfaceIpDetailParam SetL3NetworkRouterInterfaceIp详细参数
type SetL3NetworkRouterInterfaceIpDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
	rest string `json:"routerInterfaceIp" validate:"required"` // 必填
}

// SetL3NetworkRouterInterfaceIpParam SetL3NetworkRouterInterfaceIp请求参数
type SetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params SetL3NetworkRouterInterfaceIpDetailParam `json:"params"` // 详细参数
}

