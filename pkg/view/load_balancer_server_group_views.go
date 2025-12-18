// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerServerGroupInventoryView LoadBalancerServerGroup
type LoadBalancerServerGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LoadBalancerUuid string `json:"loadBalancerUuid,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ListenerServerGroupRefs []LoadBalancerListenerServerGroupRefInventoryView `json:"listenerServerGroupRefs,omitempty"`
	ServerIps []LoadBalancerServerGroupServerIpInventoryView `json:"serverIps,omitempty"`
	VmNicRefs []LoadBalancerServerGroupVmNicRefInventoryView `json:"vmNicRefs,omitempty"`
}

