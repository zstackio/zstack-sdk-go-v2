// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketInventoryView Ticket
type TicketInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	Request []TicketRequestView `json:"request,omitempty"`
	AccountSystemType string `json:"accountSystemType,omitempty"`
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	AccountSystemContext interface{} `json:"accountSystemContext,omitempty"`
	CurrentFlowUuid string `json:"currentFlowUuid,omitempty"`
	FlowCollectionUuid string `json:"flowCollectionUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// CreateTicketEventView CreateTicketEvent
type CreateTicketEventView struct {
	Inventory TicketInventoryView `json:"inventory,omitempty"`
}

// UpdateTicketRequestEventView UpdateTicketRequestEvent
type UpdateTicketRequestEventView struct {
	Inventory TicketInventoryView `json:"inventory,omitempty"`
}

// DeleteTicketEventView DeleteTicketEvent
type DeleteTicketEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryTicketView QueryTicket
type QueryTicketView struct {
	Inventories []TicketInventoryView `json:"inventories,omitempty"`
}

// ChangeTicketStatusEventView ChangeTicketStatusEvent
type ChangeTicketStatusEventView struct {
	Inventory TicketInventoryView `json:"inventory,omitempty"`
}

