// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TicketTypeInventoryView TicketType
type TicketTypeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Requests string `json:"requests,omitempty"`
	AdminOnly bool `json:"adminOnly,omitempty"`
}

// QueryTicketTypeView QueryTicketType
type QueryTicketTypeView struct {
	Inventories []TicketTypeInventoryView `json:"inventories,omitempty"`
}

