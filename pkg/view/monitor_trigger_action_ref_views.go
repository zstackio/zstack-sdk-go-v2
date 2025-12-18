// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorTriggerActionRefInventoryView MonitorTriggerActionRef
type MonitorTriggerActionRefInventoryView struct {
	rest string `json:"triggerUuid,omitempty"`
	rest string `json:"actionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

