// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MediaInventoryView Media
type MediaInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

