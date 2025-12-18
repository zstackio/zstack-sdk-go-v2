// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TicketTypeInventoryView TicketType
type TicketTypeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"requests,omitempty"`
	rest bool `json:"adminOnly,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

