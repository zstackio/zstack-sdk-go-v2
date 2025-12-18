// Copyright (c) ZStack.io, Inc.

package param

// RemoveTicketTypesFromTicketFlowCollectionDetailParam RemoveTicketTypesFromTicketFlowCollection详细参数
type RemoveTicketTypesFromTicketFlowCollectionDetailParam struct {
	rest string `json:"ticketFlowCollectionUuid" validate:"required"` // 必填
	rest []string `json:"ticketTypeUuids" validate:"required"` // 必填
}

// RemoveTicketTypesFromTicketFlowCollectionParam RemoveTicketTypesFromTicketFlowCollection请求参数
type RemoveTicketTypesFromTicketFlowCollectionParam struct {
	BaseParam
	Params RemoveTicketTypesFromTicketFlowCollectionDetailParam `json:"params"` // 详细参数
}

