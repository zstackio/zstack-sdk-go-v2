// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSSmsReceiverInventoryView SNSSmsReceiver
type SNSSmsReceiverInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

