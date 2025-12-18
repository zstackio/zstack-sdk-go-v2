// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorGroupEventSubscriptionInventoryView MonitorGroupEventSubscription
type MonitorGroupEventSubscriptionInventoryView struct {
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"eventSubscriptionUuid,omitempty"`
	rest string `json:"eventRuleTemplateUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest string `json:"uuid,omitempty"`
}

