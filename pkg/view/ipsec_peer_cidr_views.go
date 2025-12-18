// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IPsecPeerCidrInventoryView IPsecPeerCidr
type IPsecPeerCidrInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"cidr,omitempty"`
	rest string `json:"connectionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

