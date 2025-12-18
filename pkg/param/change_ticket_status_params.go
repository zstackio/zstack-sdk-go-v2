// Copyright (c) ZStack.io, Inc.

package param

// ChangeTicketStatusDetailParam ChangeTicketStatus detail param
type ChangeTicketStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StatusEvent string `json:"statusEvent" validate:"required"`
	Comment string `json:"comment,omitempty"`
}

// ChangeTicketStatusParam ChangeTicketStatus request param
type ChangeTicketStatusParam struct {
	BaseParam
	Params ChangeTicketStatusDetailParam `json:"params"`
}
