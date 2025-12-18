// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectVirtualIDGroupRefInventoryView IAM2ProjectVirtualIDGroupRef
type IAM2ProjectVirtualIDGroupRefInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

