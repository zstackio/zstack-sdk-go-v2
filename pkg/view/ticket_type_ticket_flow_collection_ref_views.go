// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketTypeTicketFlowCollectionRefInventoryView TicketTypeTicketFlowCollectionRef
type TicketTypeTicketFlowCollectionRefInventoryView struct {
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	TicketFlowCollectionUuid string `json:"ticketFlowCollectionUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

