// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ChassisOfferingInventoryView BareMetal2ChassisOffering
type BareMetal2ChassisOfferingInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest string `json:"cpuModelName,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest string `json:"bootMode,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"provisionType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

