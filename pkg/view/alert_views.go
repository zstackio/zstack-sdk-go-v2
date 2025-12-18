// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AlertInventoryView Alert
type AlertInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"triggerUuid,omitempty"`
	rest string `json:"targetResourceUuid,omitempty"`
	rest string `json:"content,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

