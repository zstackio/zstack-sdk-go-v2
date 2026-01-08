// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TicketStatusHistoryInventoryView TicketStatusHistory
type TicketStatusHistoryInventoryView struct {
	Uuid                 string      `json:"uuid,omitempty"`
	Sequence             int         `json:"sequence,omitempty"`
	TicketUuid           string      `json:"ticketUuid,omitempty"`
	FromStatus           string      `json:"fromStatus,omitempty"`
	ToStatus             string      `json:"toStatus,omitempty"`
	Comment              string      `json:"comment,omitempty"`
	OperationContextType string      `json:"operationContextType,omitempty"`
	OperationContext     interface{} `json:"operationContext,omitempty"`
	OperatorType         string      `json:"operatorType,omitempty"`
	OperatorUuid         string      `json:"operatorUuid,omitempty"`
	FlowName             string      `json:"flowName,omitempty"`
	CreateDate           time.Time   `json:"createDate,omitempty"`
	LastOpDate           time.Time   `json:"lastOpDate,omitempty"`
}

// QueryTicketHistoryView QueryTicketHistory
type QueryTicketHistoryView struct {
	Inventories []TicketStatusHistoryInventoryView `json:"inventories,omitempty"`
}
