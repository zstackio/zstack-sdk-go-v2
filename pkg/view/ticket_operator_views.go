// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketOperatorInventoryView TicketOperator
type TicketOperatorInventoryView struct {
	BaseInfoView
	BaseTimeView
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	OperatorType string `json:"operatorType,omitempty"`
	OperatorContext string `json:"operatorContext,omitempty"`
}

