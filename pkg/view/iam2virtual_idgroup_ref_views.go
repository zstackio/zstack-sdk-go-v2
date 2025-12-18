// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2VirtualIDGroupRefInventoryView IAM2VirtualIDGroupRef
type IAM2VirtualIDGroupRefInventoryView struct {
	rest string `json:"virtualIDUuid,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

