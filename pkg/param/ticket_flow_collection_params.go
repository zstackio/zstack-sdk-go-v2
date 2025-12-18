// Copyright (c) ZStack.io, Inc.

package param

// DeleteTicketFlowCollectionDetailParam DeleteTicketFlowCollection详细参数
type DeleteTicketFlowCollectionDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteTicketFlowCollectionParam DeleteTicketFlowCollection请求参数
type DeleteTicketFlowCollectionParam struct {
	BaseParam
	Params DeleteTicketFlowCollectionDetailParam `json:"params"` // 详细参数
}

