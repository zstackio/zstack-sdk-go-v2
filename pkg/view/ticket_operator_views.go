// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketOperatorInventoryView TicketOperator
type TicketOperatorInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	OperatorType string `json:"operatorType,omitempty"`
	OperatorContext string `json:"operatorContext,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

