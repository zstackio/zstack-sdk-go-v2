// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SshKeyPairInventoryView SshKeyPair
type SshKeyPairInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"publicKey,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

