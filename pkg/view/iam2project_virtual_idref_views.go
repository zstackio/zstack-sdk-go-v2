// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectVirtualIDRefInventoryView IAM2ProjectVirtualIDRef
type IAM2ProjectVirtualIDRefInventoryView struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	VirtualIDUuid string `json:"virtualIDUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

