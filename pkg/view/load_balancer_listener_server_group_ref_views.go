// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerListenerServerGroupRefInventoryView LoadBalancerListenerServerGroupRef
type LoadBalancerListenerServerGroupRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"listenerUuid,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

