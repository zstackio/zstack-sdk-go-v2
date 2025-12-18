// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerListenerVmNicRefInventoryView LoadBalancerListenerVmNicRef
type LoadBalancerListenerVmNicRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"listenerUuid,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

