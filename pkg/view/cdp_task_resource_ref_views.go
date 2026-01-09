// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// CdpTaskResourceRefInventoryView CdpTaskResourceRef
type CdpTaskResourceRefInventoryView struct {
	TaskUuid *string `json:"taskUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

