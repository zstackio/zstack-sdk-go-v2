// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NvmeLunHostRefInventoryView NvmeLunHostRef
type NvmeLunHostRefInventoryView struct {
	rest string `json:"nvmeLunUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"path,omitempty"`
	rest string `json:"hctl,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

