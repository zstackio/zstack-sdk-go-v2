// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerServerGroupVmNicRefInventoryView LoadBalancerServerGroupVmNicRef
type LoadBalancerServerGroupVmNicRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest int64 `json:"weight,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

