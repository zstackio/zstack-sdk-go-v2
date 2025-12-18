// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmUsageInventoryView VmUsage
type VmUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"name,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest int64 `json:"rootVolumeSize,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

