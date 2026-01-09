// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketFlowCollectionInventoryView TicketFlowCollection
type TicketFlowCollectionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Flows []TicketFlowInventoryView `json:"flows,omitempty"`
	TicketTypeUuids []string `json:"ticketTypeUuids,omitempty"`
}

// QueryTicketFlowCollectionView QueryTicketFlowCollection
type QueryTicketFlowCollectionView struct {
	Inventories []TicketFlowCollectionInventoryView `json:"inventories,omitempty"`
}

// DeleteTicketFlowCollectionEventView DeleteTicketFlowCollectionEvent
type DeleteTicketFlowCollectionEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeTicketFlowCollectionStateEventView ChangeTicketFlowCollectionStateEvent
type ChangeTicketFlowCollectionStateEventView struct {
	Inventory TicketFlowCollectionInventoryView `json:"inventory,omitempty"`
}

// CreateTickFlowCollectionEventView CreateTickFlowCollectionEvent
type CreateTickFlowCollectionEventView struct {
	Inventory TicketFlowCollectionInventoryView `json:"inventory,omitempty"`
}

