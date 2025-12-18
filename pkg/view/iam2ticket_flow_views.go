// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2TicketFlowInventoryView IAM2TicketFlow
type IAM2TicketFlowInventoryView struct {
	rest string `json:"approverUuid,omitempty"`
	rest bool `json:"valid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"parentFlowUuid,omitempty"`
	rest string `json:"flowContext,omitempty"`
	rest string `json:"flowContextType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"collectionUuid,omitempty"`
}

