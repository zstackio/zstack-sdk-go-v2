// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ResourceDirectoryRefInventoryView ResourceDirectoryRef
type ResourceDirectoryRefInventoryView struct {
	Id            int64     `json:"id,omitempty"`
	ResourceUuid  string    `json:"resourceUuid,omitempty"`
	DirectoryUuid string    `json:"directoryUuid,omitempty"`
	ResourceType  string    `json:"resourceType,omitempty"`
	CreateDate    time.Time `json:"createDate,omitempty"`
	LastOpDate    time.Time `json:"lastOpDate,omitempty"`
}
