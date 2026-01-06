// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ThirdClientAccountRefInventoryView ThirdClientAccountRef
type ThirdClientAccountRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ClientUuid string `json:"clientUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

