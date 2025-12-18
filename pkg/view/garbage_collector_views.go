// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GarbageCollectorInventoryView GarbageCollector
type GarbageCollectorInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"runnerClass,omitempty"`
	rest string `json:"context,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"managementNodeUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

