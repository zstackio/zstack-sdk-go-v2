// Copyright (c) ZStack.io, Inc.

package param

// DeleteTicketFlowCollectionDetailParam DeleteTicketFlowCollection detail param
type DeleteTicketFlowCollectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteTicketFlowCollectionParam DeleteTicketFlowCollection request param
type DeleteTicketFlowCollectionParam struct {
	BaseParam
	Params DeleteTicketFlowCollectionDetailParam `json:"params"`
}
