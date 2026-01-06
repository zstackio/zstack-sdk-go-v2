// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HybridConnectionRefInventoryView HybridConnectionRef
type HybridConnectionRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	ConnectionType string `json:"connectionType,omitempty"`
	ConnectionUuid string `json:"connectionUuid,omitempty"`
	Direction string `json:"direction,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

