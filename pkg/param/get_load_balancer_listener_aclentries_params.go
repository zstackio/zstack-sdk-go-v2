// Copyright (c) ZStack.io, Inc.

package param

// GetLoadBalancerListenerACLEntriesDetailParam GetLoadBalancerListenerACLEntries detail param
type GetLoadBalancerListenerACLEntriesDetailParam struct {
	ListenerUuids []string `json:"listenerUuids,omitempty"`
	Type string `json:"type,omitempty"`
}

// GetLoadBalancerListenerACLEntriesParam GetLoadBalancerListenerACLEntries request param
type GetLoadBalancerListenerACLEntriesParam struct {
	BaseParam
	Params GetLoadBalancerListenerACLEntriesDetailParam `json:"params"`
}
