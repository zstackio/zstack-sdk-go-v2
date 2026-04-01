// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2TicketFlowInventoryView IAM2TicketFlow
type IAM2TicketFlowInventoryView struct {
	BaseInfoView
	BaseTimeView
	ApproverUuid string `json:"approverUuid,omitempty"`
	Valid bool `json:"valid,omitempty"`
	Description string `json:"description,omitempty"`
	ParentFlowUuid string `json:"parentFlowUuid,omitempty"`
	FlowContext string `json:"flowContext,omitempty"`
	FlowContextType string `json:"flowContextType,omitempty"`
	CollectionUuid string `json:"collectionUuid,omitempty"`
}

// DeleteIAM2TicketFlowEventView DeleteIAM2TicketFlowEvent
type DeleteIAM2TicketFlowEventView struct {
	Success bool `json:"success,omitempty"`
}

