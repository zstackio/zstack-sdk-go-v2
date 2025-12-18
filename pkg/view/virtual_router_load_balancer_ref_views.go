// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VirtualRouterLoadBalancerRefInventoryView VirtualRouterLoadBalancerRef
type VirtualRouterLoadBalancerRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"virtualRouterVmUuid,omitempty"`
	rest string `json:"loadBalancerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

