// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MonitorTriggerInventoryView MonitorTrigger
type MonitorTriggerInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"expression,omitempty"`
	rest string `json:"recoveryExpression,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"state,omitempty"`
	rest int `json:"duration,omitempty"`
	rest string `json:"targetResourceUuid,omitempty"`
	rest time.Time `json:"lastStatusChangeTime,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

