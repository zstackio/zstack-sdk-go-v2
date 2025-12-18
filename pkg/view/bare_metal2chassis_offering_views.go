// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ChassisOfferingInventoryView BareMetal2ChassisOffering
type BareMetal2ChassisOfferingInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	CpuModelName string `json:"cpuModelName,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	BootMode string `json:"bootMode,omitempty"`
	State string `json:"state,omitempty"`
	ProvisionType string `json:"provisionType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

