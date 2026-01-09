// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IPsecL3NetworkRefInventoryView IPsecL3NetworkRef
type IPsecL3NetworkRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ConnectionUuid *string `json:"connectionUuid,omitempty"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

