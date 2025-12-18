// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmSchedHistoryInventoryView VmSchedHistory
type VmSchedHistoryInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"schedType,omitempty"`
	rest string `json:"schedReason,omitempty"`
	rest string `json:"failReason,omitempty"`
	rest bool `json:"success,omitempty"`
	rest string `json:"lastHostUuid,omitempty"`
	rest string `json:"destHostUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
}

