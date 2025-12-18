// Copyright (c) ZStack.io, Inc.

package param

// AddTicketTypesToTicketFlowCollectionDetailParam AddTicketTypesToTicketFlowCollection detail param
type AddTicketTypesToTicketFlowCollectionDetailParam struct {
	TicketFlowCollectionUuid string `json:"ticketFlowCollectionUuid" validate:"required"`
	TicketTypeUuids []string `json:"ticketTypeUuids" validate:"required"`
}

// AddTicketTypesToTicketFlowCollectionParam AddTicketTypesToTicketFlowCollection request param
type AddTicketTypesToTicketFlowCollectionParam struct {
	BaseParam
	Params AddTicketTypesToTicketFlowCollectionDetailParam `json:"params"`
}
