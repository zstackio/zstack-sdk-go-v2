// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectResourceRefInventoryView IAM2ProjectResourceRef
type IAM2ProjectResourceRefInventoryView struct {
	ProjectUuid  string    `json:"projectUuid,omitempty"`
	ResourceUuid string    `json:"resourceUuid,omitempty"`
	ResourceType string    `json:"resourceType,omitempty"`
	CreateDate   time.Time `json:"createDate,omitempty"`
	LastOpDate   time.Time `json:"lastOpDate,omitempty"`
}
