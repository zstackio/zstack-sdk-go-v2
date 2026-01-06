// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectVirtualIDGroupRefInventoryView IAM2ProjectVirtualIDGroupRef
type IAM2ProjectVirtualIDGroupRefInventoryView struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

