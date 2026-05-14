// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2DpuChassisInventoryView BareMetal2DpuChassis
type BareMetal2DpuChassisInventoryView struct {
	BaseInfoView
	BaseTimeView
	Config BareMetal2DpuChassisConfigView `json:"config,omitempty"`
	DpuHost BareMetal2DpuHostInventoryView `json:"dpuHost,omitempty"`
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

