// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EventRuleTemplateInventoryView EventRuleTemplate
type EventRuleTemplateInventoryView struct {
	Name string `json:"name,omitempty"`
	MonitorTemplateUuid string `json:"monitorTemplateUuid,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	EventName string `json:"eventName,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	Labels string `json:"labels,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// AddEventRuleTemplateEventView AddEventRuleTemplateEvent
type AddEventRuleTemplateEventView struct {
	Inventory EventRuleTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteEventRuleTemplateEventView DeleteEventRuleTemplateEvent
type DeleteEventRuleTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateEventRuleTemplateEventView UpdateEventRuleTemplateEvent
type UpdateEventRuleTemplateEventView struct {
	Inventory EventRuleTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryEventRuleTemplateView QueryEventRuleTemplate
type QueryEventRuleTemplateView struct {
	Inventories []EventRuleTemplateInventoryView `json:"inventories,omitempty"`
}

