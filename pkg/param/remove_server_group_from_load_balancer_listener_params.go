// Copyright (c) ZStack.io, Inc.

package param

// RemoveServerGroupFromLoadBalancerListenerDetailParam RemoveServerGroupFromLoadBalancerListener detail param
type RemoveServerGroupFromLoadBalancerListenerDetailParam struct {
	ServerGroupUuid string `json:"serverGroupUuid" validate:"required"`
	ListenerUuid string `json:"listenerUuid" validate:"required"`
}

// RemoveServerGroupFromLoadBalancerListenerParam RemoveServerGroupFromLoadBalancerListener request param
type RemoveServerGroupFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveServerGroupFromLoadBalancerListenerDetailParam `json:"params"`
}
