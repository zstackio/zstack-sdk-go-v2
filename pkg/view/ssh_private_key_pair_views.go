// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SshPrivateKeyPairInventoryView SshPrivateKeyPair
type SshPrivateKeyPairInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
}

