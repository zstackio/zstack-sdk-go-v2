// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketOperatorInventoryView TicketOperator
type TicketOperatorInventoryView struct {
	Uuid                string    `json:"uuid,omitempty"`
	OperatorAccountUuid string    `json:"operatorAccountUuid,omitempty"`
	OperatorType        string    `json:"operatorType,omitempty"`
	OperatorContext     string    `json:"operatorContext,omitempty"`
	CreateDate          time.Time `json:"createDate,omitempty"`
	LastOpDate          time.Time `json:"lastOpDate,omitempty"`
}
