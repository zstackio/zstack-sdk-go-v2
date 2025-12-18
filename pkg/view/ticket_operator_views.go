// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TicketOperatorInventoryView TicketOperator
type TicketOperatorInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"operatorAccountUuid,omitempty"`
	rest string `json:"operatorType,omitempty"`
	rest string `json:"operatorContext,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

