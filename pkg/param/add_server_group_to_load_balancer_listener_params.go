// Copyright (c) ZStack.io, Inc.

package param

// AddServerGroupToLoadBalancerListenerDetailParam AddServerGroupToLoadBalancerListener detail param
type AddServerGroupToLoadBalancerListenerDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// AddServerGroupToLoadBalancerListenerParam AddServerGroupToLoadBalancerListener request param
type AddServerGroupToLoadBalancerListenerParam struct {
	BaseParam
	Params AddServerGroupToLoadBalancerListenerDetailParam `json:"params"`
}
