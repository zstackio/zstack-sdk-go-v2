// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AlertInventoryView Alert
type AlertInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TriggerUuid string `json:"triggerUuid,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	Content string `json:"content,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

