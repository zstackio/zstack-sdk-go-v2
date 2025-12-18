// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SharedResourceInventoryView SharedResource
type SharedResourceInventoryView struct {
	rest string `json:"ownerAccountUuid,omitempty"`
	rest string `json:"receiverAccountUuid,omitempty"`
	rest bool `json:"toPublic,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

