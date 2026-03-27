// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TicketTypeTicketFlowCollectionRefInventoryView TicketTypeTicketFlowCollectionRef
type TicketTypeTicketFlowCollectionRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	TicketFlowCollectionUuid string `json:"ticketFlowCollectionUuid,omitempty"`
}

