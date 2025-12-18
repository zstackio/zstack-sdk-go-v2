// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSSmsReceiverInventoryView SNSSmsReceiver
type SNSSmsReceiverInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"phoneNumber,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

