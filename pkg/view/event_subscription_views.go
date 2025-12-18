// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EventSubscriptionInventoryView EventSubscription
type EventSubscriptionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"namespace,omitempty"`
	rest string `json:"eventName,omitempty"`
	rest string `json:"state,omitempty"`
	rest []EventSubscriptionActionInventoryView `json:"actions,omitempty"`
	rest []EventSubscriptionLabelInventoryView `json:"labels,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest string `json:"emergencyLevel,omitempty"`
}

