// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSEmailAddressInventoryView SNSEmailAddress
type SNSEmailAddressInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

