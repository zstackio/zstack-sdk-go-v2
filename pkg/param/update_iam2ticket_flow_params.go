// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2TicketFlowDetailParam UpdateIAM2TicketFlow detail param
type UpdateIAM2TicketFlowDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApproverUuid string `json:"approverUuid,omitempty"`
	ApproverTitle string `json:"approverTitle,omitempty"`
}

// UpdateIAM2TicketFlowParam UpdateIAM2TicketFlow request param
type UpdateIAM2TicketFlowParam struct {
	BaseParam
	Params UpdateIAM2TicketFlowDetailParam `json:"params"`
}
