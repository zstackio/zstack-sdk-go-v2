// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ScsiLunHostRefInventoryView ScsiLunHostRef
type ScsiLunHostRefInventoryView struct {
	rest string `json:"scsiLunUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"path,omitempty"`
	rest string `json:"hctl,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

