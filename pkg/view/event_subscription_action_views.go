// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EventSubscriptionActionInventoryView EventSubscriptionAction
type EventSubscriptionActionInventoryView struct {
	rest string `json:"subscriptionUuid,omitempty"`
	rest string `json:"actionType,omitempty"`
	rest string `json:"actionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

