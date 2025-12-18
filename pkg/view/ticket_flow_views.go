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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CollectionUuid string `json:"collectionUuid,omitempty"`
}

