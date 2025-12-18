// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2TicketFlowDetailParam DeleteIAM2TicketFlow detail param
type DeleteIAM2TicketFlowDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIAM2TicketFlowParam DeleteIAM2TicketFlow request param
type DeleteIAM2TicketFlowParam struct {
	BaseParam
	Params DeleteIAM2TicketFlowDetailParam `json:"params"`
}
