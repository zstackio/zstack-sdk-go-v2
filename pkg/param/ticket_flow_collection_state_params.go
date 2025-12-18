// Copyright (c) ZStack.io, Inc.

package param

// ChangeTicketFlowCollectionStateDetailParam ChangeTicketFlowCollectionState详细参数
type ChangeTicketFlowCollectionStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeTicketFlowCollectionStateParam ChangeTicketFlowCollectionState请求参数
type ChangeTicketFlowCollectionStateParam struct {
	BaseParam
	Params ChangeTicketFlowCollectionStateDetailParam `json:"params"` // 详细参数
}

