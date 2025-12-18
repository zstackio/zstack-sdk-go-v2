// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FlowCollectorInventoryView FlowCollector
type FlowCollectorInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"flowMeterUuid,omitempty"`
	rest string `json:"server,omitempty"`
	rest int64 `json:"port,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

