// Copyright (c) ZStack.io, Inc.

package param

// AddServerGroupToLoadBalancerListenerDetailParam AddServerGroupToLoadBalancerListener详细参数
type AddServerGroupToLoadBalancerListenerDetailParam struct {
	rest string `json:"serverGroupUuid" validate:"required"` // 必填
	rest string `json:"listenerUuid" validate:"required"` // 必填
}

// AddServerGroupToLoadBalancerListenerParam AddServerGroupToLoadBalancerListener请求参数
type AddServerGroupToLoadBalancerListenerParam struct {
	BaseParam
	Params AddServerGroupToLoadBalancerListenerDetailParam `json:"params"` // 详细参数
}

