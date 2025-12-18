// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2TicketFlowDetailParam UpdateIAM2TicketFlow详细参数
type UpdateIAM2TicketFlowDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"approverUuid,omitempty"`
	rest string `json:"approverTitle,omitempty"`
}

// UpdateIAM2TicketFlowParam UpdateIAM2TicketFlow请求参数
type UpdateIAM2TicketFlowParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowDetailParam `json:"params"` // 详细参数
}

