// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketTypeInventoryView TicketType
type TicketTypeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Requests *string `json:"requests,omitempty"`
	AdminOnly bool `json:"adminOnly,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryTicketTypeView QueryTicketType
type QueryTicketTypeView struct {
	Inventories []TicketTypeInventoryView `json:"inventories,omitempty"`
}

