// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDGroupRefInventoryView IAM2VirtualIDGroupRef
type IAM2VirtualIDGroupRefInventoryView struct {
	VirtualIDUuid string `json:"virtualIDUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

