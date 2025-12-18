// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LoadBalancerListenerACLRefInventoryView LoadBalancerListenerACLRef
type LoadBalancerListenerACLRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"listenerUuid,omitempty"`
	rest string `json:"serverGroupUuid,omitempty"`
	rest string `json:"aclUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

