// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ChassisInventoryView BareMetal2Chassis
type BareMetal2ChassisInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"chassisOfferingUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"powerStatus,omitempty"`
	rest string `json:"provisionType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []BareMetal2ChassisNicInventoryView `json:"chassisNics,omitempty"`
	rest []BareMetal2ChassisDiskInventoryView `json:"chassisDisks,omitempty"`
	rest BareMetal2ChassisOfferingInventoryView `json:"chassisOffering,omitempty"`
}

