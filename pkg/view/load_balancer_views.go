// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerInventoryView LoadBalancer
type LoadBalancerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Ipv6VipUuid string `json:"ipv6VipUuid,omitempty"`
	Listeners []LoadBalancerListenerInventoryView `json:"listeners,omitempty"`
}

// QueryLoadBalancerView QueryLoadBalancer
type QueryLoadBalancerView struct {
	Inventories []LoadBalancerInventoryView `json:"inventories,omitempty"`
}

// GetVpcAttachedLoadBalancerView GetVpcAttachedLoadBalancer
type GetVpcAttachedLoadBalancerView struct {
	Inventories []LoadBalancerInventoryView `json:"inventories,omitempty"`
}

// CreateLoadBalancerEventView CreateLoadBalancerEvent
type CreateLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

// UpdateLoadBalancerEventView UpdateLoadBalancerEvent
type UpdateLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

// RemoveVmNicFromLoadBalancerEventView RemoveVmNicFromLoadBalancerEvent
type RemoveVmNicFromLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

// DeleteLoadBalancerEventView DeleteLoadBalancerEvent
type DeleteLoadBalancerEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachVipToLoadBalancerEventView AttachVipToLoadBalancerEvent
type AttachVipToLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

// RefreshLoadBalancerEventView RefreshLoadBalancerEvent
type RefreshLoadBalancerEventView struct {
	Inventory LoadBalancerInventoryView `json:"inventory,omitempty"`
}

