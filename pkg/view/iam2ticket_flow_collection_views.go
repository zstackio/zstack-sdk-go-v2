// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2TicketFlowCollectionInventoryView IAM2TicketFlowCollection
type IAM2TicketFlowCollectionInventoryView struct {
	BaseInfoView
	BaseTimeView
	ProjectUuid *string `json:"projectUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Flows []TicketFlowInventoryView `json:"flows,omitempty"`
	TicketTypeUuids []string `json:"ticketTypeUuids,omitempty"`
}

