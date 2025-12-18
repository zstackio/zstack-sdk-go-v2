// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2UsageInventoryView BareMetal2Usage
type BareMetal2UsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"bareMetal2ChassisOfferingUuid,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"vmName,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
}

