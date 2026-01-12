// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ArchiveTicketInventoryView ArchiveTicket
type ArchiveTicketInventoryView struct {
	BaseInfoView
	BaseTimeView
	TicketUuid *string `json:"ticketUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	Status string `json:"status,omitempty"`
	Request []TicketRequestView `json:"request,omitempty"`
	AccountSystemType *string `json:"accountSystemType,omitempty"`
	AccountSystemContext interface{} `json:"accountSystemContext,omitempty"`
	CurrentFlowUuid *string `json:"currentFlowUuid,omitempty"`
	FlowCollectionUuid *string `json:"flowCollectionUuid,omitempty"`
	TicketTypeUuid *string `json:"ticketTypeUuid,omitempty"`
}

// QueryArchiveTicketView QueryArchiveTicket
type QueryArchiveTicketView struct {
	Inventories []ArchiveTicketInventoryView `json:"inventories,omitempty"`
}

