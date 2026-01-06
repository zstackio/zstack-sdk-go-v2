// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LoadBalancerListenerServerGroupRefInventoryView LoadBalancerListenerServerGroupRef
type LoadBalancerListenerServerGroupRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

