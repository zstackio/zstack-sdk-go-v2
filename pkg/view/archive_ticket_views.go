// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ArchiveTicketInventoryView ArchiveTicket
type ArchiveTicketInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TicketUuid string `json:"ticketUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	Status string `json:"status,omitempty"`
	Request []TicketRequestView `json:"request,omitempty"`
	AccountSystemType string `json:"accountSystemType,omitempty"`
	AccountSystemContext interface{} `json:"accountSystemContext,omitempty"`
	CurrentFlowUuid string `json:"currentFlowUuid,omitempty"`
	FlowCollectionUuid string `json:"flowCollectionUuid,omitempty"`
	TicketTypeUuid string `json:"ticketTypeUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryArchiveTicketView QueryArchiveTicket
type QueryArchiveTicketView struct {
	Inventories []ArchiveTicketInventoryView `json:"inventories,omitempty"`
}

