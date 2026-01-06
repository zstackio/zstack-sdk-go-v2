// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerListenerVmNicRefInventoryView LoadBalancerListenerVmNicRef
type LoadBalancerListenerVmNicRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

