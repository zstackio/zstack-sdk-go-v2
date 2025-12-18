// Copyright (c) ZStack.io, Inc.

package view

import "time"

// TicketStatusHistoryInventoryView TicketStatusHistory
type TicketStatusHistoryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int `json:"sequence,omitempty"`
	rest string `json:"ticketUuid,omitempty"`
	rest string `json:"fromStatus,omitempty"`
	rest string `json:"toStatus,omitempty"`
	rest string `json:"comment,omitempty"`
	rest string `json:"operationContextType,omitempty"`
	rest interface{} `json:"operationContext,omitempty"`
	rest string `json:"operatorType,omitempty"`
	rest string `json:"operatorUuid,omitempty"`
	rest string `json:"flowName,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

