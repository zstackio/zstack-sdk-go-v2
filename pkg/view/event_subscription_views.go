// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EventSubscriptionInventoryView EventSubscription
type EventSubscriptionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	EventName string `json:"eventName,omitempty"`
	State string `json:"state,omitempty"`
	Actions []EventSubscriptionActionInventoryView `json:"actions,omitempty"`
	Labels []EventSubscriptionLabelInventoryView `json:"labels,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
}

