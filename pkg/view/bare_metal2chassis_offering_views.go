// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2ChassisOfferingInventoryView BareMetal2ChassisOffering
type BareMetal2ChassisOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	CpuModelName string `json:"cpuModelName,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	BootMode string `json:"bootMode,omitempty"`
	State string `json:"state,omitempty"`
	ProvisionType string `json:"provisionType,omitempty"`
}

// UpdateBareMetal2ChassisOfferingEventView UpdateBareMetal2ChassisOfferingEvent
type UpdateBareMetal2ChassisOfferingEventView struct {
	Inventory BareMetal2ChassisOfferingInventoryView `json:"inventory,omitempty"`
}

// ChangeBareMetal2ChassisOfferingStateEventView ChangeBareMetal2ChassisOfferingStateEvent
type ChangeBareMetal2ChassisOfferingStateEventView struct {
	Inventory BareMetal2ChassisOfferingInventoryView `json:"inventory,omitempty"`
}

// QueryBareMetal2ChassisOfferingView QueryBareMetal2ChassisOffering
type QueryBareMetal2ChassisOfferingView struct {
	Inventories []BareMetal2ChassisOfferingInventoryView `json:"inventories,omitempty"`
}

