// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BaremetalChassisInventoryView BaremetalChassis
type BaremetalChassisInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	PxeServerUuid *string `json:"pxeServerUuid,omitempty"`
	IpmiAddress *string `json:"ipmiAddress,omitempty"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
	IpmiUsername *string `json:"ipmiUsername,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	HardwareInfos []BaremetalHardwareInfoInventoryView `json:"hardwareInfos,omitempty"`
}

// QueryBaremetalChassisView QueryBaremetalChassis
type QueryBaremetalChassisView struct {
	Inventories []BaremetalChassisInventoryView `json:"inventories,omitempty"`
}

// InspectBaremetalChassisEventView InspectBaremetalChassisEvent
type InspectBaremetalChassisEventView struct {
	Inventory BaremetalChassisInventoryView `json:"inventory,omitempty"`
}

// UpdateBaremetalChassisEventView UpdateBaremetalChassisEvent
type UpdateBaremetalChassisEventView struct {
	Inventory BaremetalChassisInventoryView `json:"inventory,omitempty"`
}

// ChangeBaremetalChassisStateEventView ChangeBaremetalChassisStateEvent
type ChangeBaremetalChassisStateEventView struct {
	Inventory BaremetalChassisInventoryView `json:"inventory,omitempty"`
}

// DeleteBaremetalChassisEventView DeleteBaremetalChassisEvent
type DeleteBaremetalChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateBaremetalChassisEventView CreateBaremetalChassisEvent
type CreateBaremetalChassisEventView struct {
	Inventory BaremetalChassisInventoryView `json:"inventory,omitempty"`
}

