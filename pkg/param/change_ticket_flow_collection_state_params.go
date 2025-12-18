// Copyright (c) ZStack.io, Inc.

package param

// ChangeTicketFlowCollectionStateDetailParam ChangeTicketFlowCollectionState detail param
type ChangeTicketFlowCollectionStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeTicketFlowCollectionStateParam ChangeTicketFlowCollectionState request param
type ChangeTicketFlowCollectionStateParam struct {
	BaseParam
	Params ChangeTicketFlowCollectionStateDetailParam `json:"params"`
}
