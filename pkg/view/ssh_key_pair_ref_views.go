// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SshKeyPairRefInventoryView SshKeyPairRef
type SshKeyPairRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	SshKeyPairUuid string `json:"sshKeyPairUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

