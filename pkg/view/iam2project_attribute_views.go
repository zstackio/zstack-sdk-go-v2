// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectAttributeInventoryView IAM2ProjectAttribute
type IAM2ProjectAttributeInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"value,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

