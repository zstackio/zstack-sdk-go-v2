// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SshKeyPairInventoryView SshKeyPair
type SshKeyPairInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// CreateSshKeyPairEventView CreateSshKeyPairEvent
type CreateSshKeyPairEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// GenerateSshKeyPairView GenerateSshKeyPair
type GenerateSshKeyPairView struct {
	Inventory SshPrivateKeyPairInventoryView `json:"inventory,omitempty"`
}

// UpdateSshKeyPairEventView UpdateSshKeyPairEvent
type UpdateSshKeyPairEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// AttachSshKeyPairToVmInstanceEventView AttachSshKeyPairToVmInstanceEvent
type AttachSshKeyPairToVmInstanceEventView struct {
	Inventory SshKeyPairInventoryView `json:"inventory,omitempty"`
}

// QuerySshKeyPairView QuerySshKeyPair
type QuerySshKeyPairView struct {
	Inventories []SshKeyPairInventoryView `json:"inventories,omitempty"`
}

// DeleteSshKeyPairEventView DeleteSshKeyPairEvent
type DeleteSshKeyPairEventView struct {
	Success bool `json:"success,omitempty"`
}

