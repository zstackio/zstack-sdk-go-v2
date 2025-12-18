// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSEmailAddressInventoryView SNSEmailAddress
type SNSEmailAddressInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"emailAddress,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

