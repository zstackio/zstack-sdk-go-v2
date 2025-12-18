// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IPsecL3NetworkRefInventoryView IPsecL3NetworkRef
type IPsecL3NetworkRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"connectionUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

