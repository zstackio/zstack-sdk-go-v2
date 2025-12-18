// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TicketInventoryView Ticket
type TicketInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"status,omitempty"`
	rest []interface{} `json:"request,omitempty"`
	rest string `json:"accountSystemType,omitempty"`
	rest string `json:"ticketTypeUuid,omitempty"`
	rest interface{} `json:"accountSystemContext,omitempty"`
	rest string `json:"currentFlowUuid,omitempty"`
	rest string `json:"flowCollectionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

