// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerInventoryView LoadBalancer
type LoadBalancerInventoryView struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Ipv6VipUuid string `json:"ipv6VipUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Listeners []LoadBalancerListenerInventoryView `json:"listeners,omitempty"`
}

