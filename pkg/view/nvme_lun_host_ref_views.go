// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NvmeLunHostRefInventoryView NvmeLunHostRef
type NvmeLunHostRefInventoryView struct {
	NvmeLunUuid string `json:"nvmeLunUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Path string `json:"path,omitempty"`
	Hctl string `json:"hctl,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

