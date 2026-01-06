// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorGroupInstanceInventoryView MonitorGroupInstance
type MonitorGroupInstanceInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	InstanceResourceType string `json:"instanceResourceType,omitempty"`
	InstanceUuid string `json:"instanceUuid,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// QueryMonitorGroupInstanceView QueryMonitorGroupInstance
type QueryMonitorGroupInstanceView struct {
	Inventories []MonitorGroupInstanceInventoryView `json:"inventories,omitempty"`
}

// AddInstanceToMonitorGroupEventView AddInstanceToMonitorGroupEvent
type AddInstanceToMonitorGroupEventView struct {
	Inventory MonitorGroupInstanceInventoryView `json:"inventory,omitempty"`
}

