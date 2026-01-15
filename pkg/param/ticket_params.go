// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateTicketParamDetail CreateTicket detail param
type CreateTicketParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Requests []TicketRequestParam `json:"requests" validate:"required"`
	FlowCollectionUuid string `json:"flowCollectionUuid,omitempty"`
	AccountSystemType string `json:"accountSystemType" validate:"required"`
	AccountSystemContext interface{} `json:"accountSystemContext" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateTicketParam CreateTicket request param
type CreateTicketParam struct {
	BaseParam
	CreateTicket CreateTicketParamDetail `json:"createTicket"`
}
// DeleteTicketParamDetail DeleteTicket detail param
type DeleteTicketParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteTicketParam DeleteTicket request param
type DeleteTicketParam struct {
	BaseParam
	DeleteTicket DeleteTicketParamDetail `json:"deleteTicket"`
}
