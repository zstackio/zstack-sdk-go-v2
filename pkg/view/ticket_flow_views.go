// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketFlowInventoryView TicketFlow
type TicketFlowInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ParentFlowUuid string `json:"parentFlowUuid,omitempty"`
	FlowContext string `json:"flowContext,omitempty"`
	FlowContextType string `json:"flowContextType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	CollectionUuid string `json:"collectionUuid,omitempty"`
}

// UpdateIAM2TicketFlowEventView UpdateIAM2TicketFlowEvent
type UpdateIAM2TicketFlowEventView struct {
	Inventory TicketFlowInventoryView `json:"inventory,omitempty"`
}

// QueryTicketFlowView QueryTicketFlow
type QueryTicketFlowView struct {
	Inventories []TicketFlowInventoryView `json:"inventories,omitempty"`
}

// AddIAM2TicketFlowEventView AddIAM2TicketFlowEvent
type AddIAM2TicketFlowEventView struct {
	Inventory TicketFlowInventoryView `json:"inventory,omitempty"`
}

