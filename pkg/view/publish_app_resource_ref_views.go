// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PublishAppResourceRefInventoryView PublishAppResourceRef
type PublishAppResourceRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	AppUuid string `json:"appUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

