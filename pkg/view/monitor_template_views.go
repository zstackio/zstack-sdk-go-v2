// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorTemplateInventoryView MonitorTemplate
type MonitorTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	MonitorGroupTemplateRefs []MonitorGroupTemplateRefVOView `json:"monitorGroupTemplateRefs,omitempty"`
}

// UpdateMonitorTemplateEventView UpdateMonitorTemplateEvent
type UpdateMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// CloneMonitorTemplateEventView CloneMonitorTemplateEvent
type CloneMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryMonitorTemplateView QueryMonitorTemplate
type QueryMonitorTemplateView struct {
	Inventories []MonitorTemplateInventoryView `json:"inventories,omitempty"`
}

// CreateMonitorTemplateEventView CreateMonitorTemplateEvent
type CreateMonitorTemplateEventView struct {
	Inventory MonitorTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteMonitorTemplateEventView DeleteMonitorTemplateEvent
type DeleteMonitorTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

