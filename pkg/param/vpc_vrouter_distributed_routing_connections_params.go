// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterDistributedRoutingConnectionsDetailParam GetVpcVRouterDistributedRoutingConnections详细参数
type GetVpcVRouterDistributedRoutingConnectionsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVpcVRouterDistributedRoutingConnectionsParam GetVpcVRouterDistributedRoutingConnections请求参数
type GetVpcVRouterDistributedRoutingConnectionsParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingConnectionsDetailParam `json:"params"` // 详细参数
}

