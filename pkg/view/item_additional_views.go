// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ItemInventoryView Item
type ItemInventoryView struct {
	Name string `json:"name,omitempty"`
	ReadableName *string `json:"readableName,omitempty"`
}

