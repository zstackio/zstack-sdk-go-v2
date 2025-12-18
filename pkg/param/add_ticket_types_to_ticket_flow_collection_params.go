// Copyright (c) ZStack.io, Inc.

package param

// AddTicketTypesToTicketFlowCollectionDetailParam AddTicketTypesToTicketFlowCollection详细参数
type AddTicketTypesToTicketFlowCollectionDetailParam struct {
	rest string `json:"ticketFlowCollectionUuid" validate:"required"` // 必填
	rest []string `json:"ticketTypeUuids" validate:"required"` // 必填
}

// AddTicketTypesToTicketFlowCollectionParam AddTicketTypesToTicketFlowCollection请求参数
type AddTicketTypesToTicketFlowCollectionParam struct {
	BaseParam
	Params AddTicketTypesToTicketFlowCollectionDetailParam `json:"params"` // 详细参数
}

