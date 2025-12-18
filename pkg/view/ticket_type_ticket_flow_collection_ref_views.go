// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TicketTypeTicketFlowCollectionRefInventoryView TicketTypeTicketFlowCollectionRef
type TicketTypeTicketFlowCollectionRefInventoryView struct {
	rest string `json:"ticketTypeUuid,omitempty"`
	rest string `json:"ticketFlowCollectionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

