// Copyright (c) ZStack.io, Inc.

package view

// GetLoadBalancerOwnerView GetLoadBalancerOwner
type GetLoadBalancerOwnerView struct {
	Type string `json:"type,omitempty"`
	Vpc VpcRouterVmInventoryView `json:"vpc,omitempty"`
	VpcHa VpcHaGroupInventoryView `json:"vpcHa,omitempty"`
	Slb SlbGroupInventoryView `json:"slb,omitempty"`
	Success bool `json:"success,omitempty"`
}

