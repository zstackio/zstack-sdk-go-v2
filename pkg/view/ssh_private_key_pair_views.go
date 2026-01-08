// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SshPrivateKeyPairInventoryView SshPrivateKeyPair
type SshPrivateKeyPairInventoryView struct {
	BaseInfoView
	BaseTimeView
	PublicKey  string `json:"publicKey,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
}
