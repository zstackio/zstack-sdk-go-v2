// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AccountResourceRefInventoryView AccountResourceRef
type AccountResourceRefInventoryView struct {
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"concreteResourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

