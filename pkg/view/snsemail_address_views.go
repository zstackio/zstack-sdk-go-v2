// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSEmailAddressInventoryView SNSEmailAddress
type SNSEmailAddressInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// AddEmailAddressToSNSEmailEndpointEventView AddEmailAddressToSNSEmailEndpointEvent
type AddEmailAddressToSNSEmailEndpointEventView struct {
	Inventory SNSEmailAddressInventoryView `json:"inventory,omitempty"`
}

// UpdateEmailAddressOfSNSEmailEndpointEventView UpdateEmailAddressOfSNSEmailEndpointEvent
type UpdateEmailAddressOfSNSEmailEndpointEventView struct {
	Inventory SNSEmailAddressInventoryView `json:"inventory,omitempty"`
}

// QuerySNSEmailAddressView QuerySNSEmailAddress
type QuerySNSEmailAddressView struct {
	Inventories []SNSEmailAddressInventoryView `json:"inventories,omitempty"`
}

