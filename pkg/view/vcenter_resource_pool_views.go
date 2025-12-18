// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VCenterResourcePoolInventoryView VCenterResourcePool
type VCenterResourcePoolInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vCenterClusterUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"morVal,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest int64 `json:"CPULimit,omitempty"`
	rest int64 `json:"CPUOverheadLimit,omitempty"`
	rest int64 `json:"CPUReservation,omitempty"`
	rest int64 `json:"CPUShares,omitempty"`
	rest string `json:"CPULevel,omitempty"`
	rest int64 `json:"memoryLimit,omitempty"`
	rest int64 `json:"memoryOverheadLimit,omitempty"`
	rest int64 `json:"memoryReservation,omitempty"`
	rest int64 `json:"memoryShares,omitempty"`
	rest string `json:"memoryLevel,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VCenterResourcePoolUsageInventoryView `json:"subResources,omitempty"`
}

