// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupInventoryView MonitorGroup
type MonitorGroupInventoryView struct {
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Actions string `json:"actions,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	MonitorGroupTemplateRefs []MonitorGroupTemplateRefVOView `json:"monitorGroupTemplateRefs,omitempty"`
}

// CreateMonitorGroupEventView CreateMonitorGroupEvent
type CreateMonitorGroupEventView struct {
	Inventory MonitorGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteMonitorGroupEventView DeleteMonitorGroupEvent
type DeleteMonitorGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryMonitorGroupView QueryMonitorGroup
type QueryMonitorGroupView struct {
	Inventories []MonitorGroupInventoryView `json:"inventories,omitempty"`
}

// UpdateMonitorGroupEventView UpdateMonitorGroupEvent
type UpdateMonitorGroupEventView struct {
	Inventory MonitorGroupInventoryView `json:"inventory,omitempty"`
}

