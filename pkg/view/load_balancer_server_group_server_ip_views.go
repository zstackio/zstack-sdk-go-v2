// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerServerGroupServerIpInventoryView LoadBalancerServerGroupServerIp
type LoadBalancerServerGroupServerIpInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest string `json:"ipAddress,omitempty"`
	rest int64 `json:"weight,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

