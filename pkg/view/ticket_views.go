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
	Request []interface{} `json:"request,omitempty"`
	AccountSystemType string `json:"accountSystemType,omitempty"`
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	AccountSystemContext interface{} `json:"accountSystemContext,omitempty"`
	CurrentFlowUuid string `json:"currentFlowUuid,omitempty"`
	FlowCollectionUuid string `json:"flowCollectionUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

