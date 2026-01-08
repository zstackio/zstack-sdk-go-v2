// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupEventSubscriptionInventoryView MonitorGroupEventSubscription
type MonitorGroupEventSubscriptionInventoryView struct {
	GroupUuid             string    `json:"groupUuid,omitempty"`
	EventSubscriptionUuid string    `json:"eventSubscriptionUuid,omitempty"`
	EventRuleTemplateUuid string    `json:"eventRuleTemplateUuid,omitempty"`
	CreateDate            time.Time `json:"createDate,omitempty"`
	Uuid                  string    `json:"uuid,omitempty"`
}

// QueryMonitorGroupEventSubscriptionView QueryMonitorGroupEventSubscription
type QueryMonitorGroupEventSubscriptionView struct {
	Inventories []MonitorGroupEventSubscriptionInventoryView `json:"inventories,omitempty"`
}
