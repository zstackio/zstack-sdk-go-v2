// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketFlowInventoryView TicketFlow
type TicketFlowInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ParentFlowUuid string `json:"parentFlowUuid,omitempty"`
	FlowContext string `json:"flowContext,omitempty"`
	FlowContextType string `json:"flowContextType,omitempty"`
	CollectionUuid string `json:"collectionUuid,omitempty"`
}

// QueryTicketFlowView QueryTicketFlow
type QueryTicketFlowView struct {
	Inventories []TicketFlowInventoryView `json:"inventories,omitempty"`
}

