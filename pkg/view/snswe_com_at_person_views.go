// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSWeComAtPersonInventoryView SNSWeComAtPerson
type SNSWeComAtPersonInventoryView struct {
	Uuid         string    `json:"uuid,omitempty"`
	UserId       string    `json:"userId,omitempty"`
	EndpointUuid string    `json:"endpointUuid,omitempty"`
	CreateDate   time.Time `json:"createDate,omitempty"`
	LastOpDate   time.Time `json:"lastOpDate,omitempty"`
	Remark       string    `json:"remark,omitempty"`
}

// AddSNSWeComAtPersonEventView AddSNSWeComAtPersonEvent
type AddSNSWeComAtPersonEventView struct {
	Inventory SNSWeComAtPersonInventoryView `json:"inventory,omitempty"`
}

// QuerySNSWeComAtPersonView QuerySNSWeComAtPerson
type QuerySNSWeComAtPersonView struct {
	Inventories []SNSWeComAtPersonInventoryView `json:"inventories,omitempty"`
}

// RemoveSNSWeComAtPersonEventView RemoveSNSWeComAtPersonEvent
type RemoveSNSWeComAtPersonEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateAtPersonOfWeComEndpointEventView UpdateAtPersonOfWeComEndpointEvent
type UpdateAtPersonOfWeComEndpointEventView struct {
	Inventory SNSWeComAtPersonInventoryView `json:"inventory,omitempty"`
}
