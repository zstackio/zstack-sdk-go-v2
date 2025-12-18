// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerInventoryView LoadBalancer
type LoadBalancerInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"ipv6VipUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []LoadBalancerListenerInventoryView `json:"listeners,omitempty"`
}

