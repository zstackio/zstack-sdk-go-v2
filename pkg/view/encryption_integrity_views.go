// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EncryptionIntegrityInventoryView EncryptionIntegrity
type EncryptionIntegrityInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	SignedText string `json:"signedText,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

