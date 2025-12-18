// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisInventoryView BareMetal2Chassis
type BareMetal2ChassisInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ChassisOfferingUuid string `json:"chassisOfferingUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	PowerStatus string `json:"powerStatus,omitempty"`
	ProvisionType string `json:"provisionType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ChassisNics []BareMetal2ChassisNicInventoryView `json:"chassisNics,omitempty"`
	ChassisDisks []BareMetal2ChassisDiskInventoryView `json:"chassisDisks,omitempty"`
	ChassisOffering BareMetal2ChassisOfferingInventoryView `json:"chassisOffering,omitempty"`
}

