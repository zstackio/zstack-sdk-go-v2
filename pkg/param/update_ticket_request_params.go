// Copyright (c) ZStack.io, Inc.

package param

// UpdateTicketRequestDetailParam UpdateTicketRequest detail param
type UpdateTicketRequestDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Requests []TicketRequestParam `json:"requests" validate:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateTicketRequestParam UpdateTicketRequest request param
type UpdateTicketRequestParam struct {
	BaseParam
	Params UpdateTicketRequestDetailParam `json:"params"`
}
