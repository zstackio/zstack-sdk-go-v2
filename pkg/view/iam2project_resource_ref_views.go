// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectResourceRefInventoryView IAM2ProjectResourceRef
type IAM2ProjectResourceRefInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

