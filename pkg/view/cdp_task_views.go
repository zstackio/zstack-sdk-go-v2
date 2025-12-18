// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CdpTaskInventoryView CdpTask
type CdpTaskInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"policyUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"taskType,omitempty"`
	rest int64 `json:"backupBandwidth,omitempty"`
	rest int64 `json:"maxCapacity,omitempty"`
	rest int64 `json:"usedCapacity,omitempty"`
	rest int64 `json:"maxLatency,omitempty"`
	rest int64 `json:"lastLatency,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []CdpTaskResourceRefInventoryView `json:"resourceRefs,omitempty"`
}

