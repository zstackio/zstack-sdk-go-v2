// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IPsecPeerCidrInventoryView IPsecPeerCidr
type IPsecPeerCidrInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
	ConnectionUuid *string `json:"connectionUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

