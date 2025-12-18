// Copyright (c) ZStack.io, Inc.

package param

// RemoveTicketTypesFromTicketFlowCollectionDetailParam RemoveTicketTypesFromTicketFlowCollection detail param
type RemoveTicketTypesFromTicketFlowCollectionDetailParam struct {
	TicketFlowCollectionUuid string `json:"ticketFlowCollectionUuid" validate:"required"`
	TicketTypeUuids []string `json:"ticketTypeUuids" validate:"required"`
}

// RemoveTicketTypesFromTicketFlowCollectionParam RemoveTicketTypesFromTicketFlowCollection request param
type RemoveTicketTypesFromTicketFlowCollectionParam struct {
	BaseParam
	Params RemoveTicketTypesFromTicketFlowCollectionDetailParam `json:"params"`
}
