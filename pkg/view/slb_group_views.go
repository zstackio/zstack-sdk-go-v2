// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SlbGroupInventoryView SlbGroup
type SlbGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	BackendType string `json:"backendType,omitempty"`
	DeployType string `json:"deployType,omitempty"`
	SlbOfferingUuid string `json:"slbOfferingUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ConfigVersion int64 `json:"configVersion,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	SlbVms []SlbVmInstanceInventoryView `json:"slbVms,omitempty"`
	Lbs []SlbLoadBalancerInventoryView `json:"lbs,omitempty"`
	Networks []SlbGroupL3NetworkRefInventoryView `json:"networks,omitempty"`
	MonitorIps []SlbGroupMonitorIpInventoryView `json:"monitorIps,omitempty"`
}

// ChangeSlbGroupMonitorIpsEventView ChangeSlbGroupMonitorIpsEvent
type ChangeSlbGroupMonitorIpsEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

// QuerySlbGroupView QuerySlbGroup
type QuerySlbGroupView struct {
	Inventories []SlbGroupInventoryView `json:"inventories,omitempty"`
}

// CreateSlbGroupEventView CreateSlbGroupEvent
type CreateSlbGroupEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteSlbGroupEventView DeleteSlbGroupEvent
type DeleteSlbGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateSlbGroupEventView UpdateSlbGroupEvent
type UpdateSlbGroupEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

// ChangeSlbGroupDeployTypeEventView ChangeSlbGroupDeployTypeEvent
type ChangeSlbGroupDeployTypeEventView struct {
	Inventory SlbGroupInventoryView `json:"inventory,omitempty"`
}

