// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResNotifySubscriptionInventoryView ResNotifySubscription
type ResNotifySubscriptionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ResourceTypes string `json:"resourceTypes,omitempty"`
	EventTypes string `json:"eventTypes,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	WebhookRef ResNotifyWebhookRefInventoryView `json:"webhookRef,omitempty"`
}

// DeleteResNotifySubscriptionEventView DeleteResNotifySubscriptionEvent
type DeleteResNotifySubscriptionEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryResNotifySubscriptionView QueryResNotifySubscription
type QueryResNotifySubscriptionView struct {
	Inventories []ResNotifySubscriptionInventoryView `json:"inventories,omitempty"`
}

// UpdateResNotifySubscriptionEventView UpdateResNotifySubscriptionEvent
type UpdateResNotifySubscriptionEventView struct {
	Inventory ResNotifySubscriptionInventoryView `json:"inventory,omitempty"`
}

