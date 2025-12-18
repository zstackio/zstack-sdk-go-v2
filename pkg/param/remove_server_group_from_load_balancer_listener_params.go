// Copyright (c) ZStack.io, Inc.

package param

// RemoveServerGroupFromLoadBalancerListenerDetailParam RemoveServerGroupFromLoadBalancerListener详细参数
type RemoveServerGroupFromLoadBalancerListenerDetailParam struct {
	rest string `json:"serverGroupUuid" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// RemoveServerGroupFromLoadBalancerListenerParam RemoveServerGroupFromLoadBalancerListener请求参数
type RemoveServerGroupFromLoadBalancerListenerParam struct {
	BaseParam
	Params RemoveServerGroupFromLoadBalancerListenerDetailParam `json:"params"` // 详细参数
}

