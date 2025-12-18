// Copyright (c) ZStack.io, Inc.

package param

// GetL3NetworkRouterInterfaceIpDetailParam GetL3NetworkRouterInterfaceIp详细参数
type GetL3NetworkRouterInterfaceIpDetailParam struct {
	rest string `json:"l3NetworkUuid" validate:"required"` // 必填
}

// GetL3NetworkRouterInterfaceIpParam GetL3NetworkRouterInterfaceIp请求参数
type GetL3NetworkRouterInterfaceIpParam struct {
	BaseParam
	Params GetL3NetworkRouterInterfaceIpDetailParam `json:"params"` // 详细参数
}

