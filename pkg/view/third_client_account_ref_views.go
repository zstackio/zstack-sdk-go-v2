// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ThirdClientAccountRefInventoryView ThirdClientAccountRef
type ThirdClientAccountRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"clientUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

