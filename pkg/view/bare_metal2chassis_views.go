// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisInventoryView BareMetal2Chassis
type BareMetal2ChassisInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	ChassisOfferingUuid *string `json:"chassisOfferingUuid,omitempty"`
	Type *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	PowerStatus *string `json:"powerStatus,omitempty"`
	ProvisionType *string `json:"provisionType,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	ChassisNics []BareMetal2ChassisNicInventoryView `json:"chassisNics,omitempty"`
	ChassisDisks []BareMetal2ChassisDiskInventoryView `json:"chassisDisks,omitempty"`
	ChassisOffering BareMetal2ChassisOfferingInventoryView `json:"chassisOffering,omitempty"`
}

// InspectBareMetal2ChassisEventView InspectBareMetal2ChassisEvent
type InspectBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// UpdateBareMetal2ChassisEventView UpdateBareMetal2ChassisEvent
type UpdateBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// QueryBareMetal2ChassisView QueryBareMetal2Chassis
type QueryBareMetal2ChassisView struct {
	Inventories []BareMetal2ChassisInventoryView `json:"inventories,omitempty"`
}

// ChangeBareMetal2ChassisStateEventView ChangeBareMetal2ChassisStateEvent
type ChangeBareMetal2ChassisStateEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// DeleteBareMetal2ChassisEventView DeleteBareMetal2ChassisEvent
type DeleteBareMetal2ChassisEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddBareMetal2ChassisEventView AddBareMetal2ChassisEvent
type AddBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// InspectBareMetal2ChassisByInstanceEventView InspectBareMetal2ChassisByInstanceEvent
type InspectBareMetal2ChassisByInstanceEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

