// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorTriggerActionInventoryView MonitorTriggerAction
type MonitorTriggerActionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"type,omitempty"`
	rest []string `json:"triggerUuids,omitempty"`
}

