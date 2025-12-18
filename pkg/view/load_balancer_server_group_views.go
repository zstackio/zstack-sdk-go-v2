// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerServerGroupInventoryView LoadBalancerServerGroup
type LoadBalancerServerGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"loadBalancerUuid,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []LoadBalancerListenerServerGroupRefInventoryView `json:"listenerServerGroupRefs,omitempty"`
	rest []LoadBalancerServerGroupServerIpInventoryView `json:"serverIps,omitempty"`
	rest []LoadBalancerServerGroupVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
}

