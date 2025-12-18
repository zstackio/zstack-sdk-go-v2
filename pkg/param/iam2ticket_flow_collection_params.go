// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2TicketFlowCollectionDetailParam UpdateIAM2TicketFlowCollection详细参数
type UpdateIAM2TicketFlowCollectionDetailParam struct {
	rest []interface{} `json:"flows,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"isDefault,omitempty"`
}

// UpdateIAM2TicketFlowCollectionParam UpdateIAM2TicketFlowCollection请求参数
type UpdateIAM2TicketFlowCollectionParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowCollectionDetailParam `json:"params"` // 详细参数
}

