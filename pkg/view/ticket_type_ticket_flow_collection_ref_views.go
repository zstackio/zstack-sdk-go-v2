// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketTypeTicketFlowCollectionRefInventoryView TicketTypeTicketFlowCollectionRef
type TicketTypeTicketFlowCollectionRefInventoryView struct {
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	TicketFlowCollectionUuid string `json:"ticketFlowCollectionUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

