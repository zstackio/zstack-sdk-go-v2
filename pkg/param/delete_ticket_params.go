// Copyright (c) ZStack.io, Inc.

package param

// DeleteTicketDetailParam DeleteTicket detail param
type DeleteTicketDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteTicketParam DeleteTicket request param
type DeleteTicketParam struct {
	BaseParam
	Params DeleteTicketDetailParam `json:"params"`
}
