// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MonitorTriggerActionRefInventoryView MonitorTriggerActionRef
type MonitorTriggerActionRefInventoryView struct {
	TriggerUuid string    `json:"triggerUuid,omitempty"`
	ActionUuid  string    `json:"actionUuid,omitempty"`
	CreateDate  time.Time `json:"createDate,omitempty"`
	LastOpDate  time.Time `json:"lastOpDate,omitempty"`
}
