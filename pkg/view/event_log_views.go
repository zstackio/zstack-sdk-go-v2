// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EventLogInventoryView EventLog
type EventLogInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"content,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"category,omitempty"`
	rest string `json:"trackingId,omitempty"`
	rest string `json:"type,omitempty"`
	rest int64 `json:"time,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

