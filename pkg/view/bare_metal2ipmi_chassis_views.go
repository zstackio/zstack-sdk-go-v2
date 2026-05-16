// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2IpmiChassisInventoryView BareMetal2IpmiChassis
type BareMetal2IpmiChassisInventoryView struct {
	BaseInfoView
	BaseTimeView
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	PowerStatus string `json:"powerStatus,omitempty"`
	ProvisionType string `json:"provisionType,omitempty"`
	ChassisNics []BareMetal2ChassisNicInventoryView `json:"chassisNics,omitempty"`
	ChassisDisks []BareMetal2ChassisDiskInventoryView `json:"chassisDisks,omitempty"`
	ChassisOffering BareMetal2ChassisOfferingInventoryView `json:"chassisOffering,omitempty"`
}

// AddBareMetal2ChassisEventView AddBareMetal2ChassisEvent
type AddBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

// UpdateBareMetal2ChassisEventView UpdateBareMetal2ChassisEvent
type UpdateBareMetal2ChassisEventView struct {
	Inventory BareMetal2ChassisInventoryView `json:"inventory,omitempty"`
}

