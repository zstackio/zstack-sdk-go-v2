// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CephPrimaryStorageMonInventoryView CephPrimaryStorageMon
type CephPrimaryStorageMonInventoryView struct {
	rest string `json:"hostname,omitempty"`
	rest int `json:"monPort,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"monAddr,omitempty"`
	rest string `json:"sshUsername,omitempty"`
	rest string `json:"sshPassword,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"monUuid,omitempty"`
}

