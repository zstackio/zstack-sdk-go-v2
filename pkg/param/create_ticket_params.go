// Copyright (c) ZStack.io, Inc.

package param

// CreateTicketDetailParam CreateTicket detail param
type CreateTicketDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Requests []interface{} `json:"requests" validate:"required"`
	FlowCollectionUuid string `json:"flowCollectionUuid,omitempty"`
	AccountSystemType string `json:"accountSystemType" validate:"required"`
	AccountSystemContext interface{} `json:"accountSystemContext" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateTicketParam CreateTicket request param
type CreateTicketParam struct {
	BaseParam
	Params CreateTicketDetailParam `json:"params"`
}
