// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EncryptionIntegrityInventoryView EncryptionIntegrity
type EncryptionIntegrityInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"signedText,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

