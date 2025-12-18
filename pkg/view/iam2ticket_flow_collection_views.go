// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2TicketFlowCollectionInventoryView IAM2TicketFlowCollection
type IAM2TicketFlowCollectionInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest bool `json:"isDefault,omitempty"`
	rest []TicketFlowInventoryView `json:"flows,omitempty"`
	rest []string `json:"ticketTypeUuids,omitempty"`
}

