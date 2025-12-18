// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectVirtualIDRefInventoryView IAM2ProjectVirtualIDRef
type IAM2ProjectVirtualIDRefInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"virtualIDUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

