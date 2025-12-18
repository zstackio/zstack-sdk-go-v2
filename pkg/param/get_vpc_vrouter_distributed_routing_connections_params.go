// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVRouterDistributedRoutingConnectionsDetailParam GetVpcVRouterDistributedRoutingConnections detail param
type GetVpcVRouterDistributedRoutingConnectionsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVpcVRouterDistributedRoutingConnectionsParam GetVpcVRouterDistributedRoutingConnections request param
type GetVpcVRouterDistributedRoutingConnectionsParam struct {
	BaseParam
	Params GetVpcVRouterDistributedRoutingConnectionsDetailParam `json:"params"`
}
