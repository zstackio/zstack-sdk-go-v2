// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2TicketFlowCollectionInventoryView IAM2TicketFlowCollection
type IAM2TicketFlowCollectionInventoryView struct {
	ProjectUuid *string `json:"projectUuid,omitempty"`
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

// UpdateTicketFlowCollectionEventView UpdateTicketFlowCollectionEvent
type UpdateTicketFlowCollectionEventView struct {
	Inventory TicketFlowCollectionInventoryView `json:"inventory,omitempty"`
}

