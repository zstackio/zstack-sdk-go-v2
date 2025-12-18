// Copyright (c) ZStack.io, Inc.

package param

// GetLoadBalancerListenerACLEntriesDetailParam GetLoadBalancerListenerACLEntries详细参数
type GetLoadBalancerListenerACLEntriesDetailParam struct {
	rest []string `json:"listenerUuids,omitempty"`
	rest string `json:"type,omitempty"`
}

// GetLoadBalancerListenerACLEntriesParam GetLoadBalancerListenerACLEntries请求参数
type GetLoadBalancerListenerACLEntriesParam struct {
	BaseParam
	Params GetLoadBalancerListenerACLEntriesDetailParam `json:"params"` // 详细参数
}

