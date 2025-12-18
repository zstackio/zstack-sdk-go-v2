// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectAccountRefInventoryView IAM2ProjectAccountRef
type IAM2ProjectAccountRefInventoryView struct {
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

