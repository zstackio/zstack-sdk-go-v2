// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RemoteVtepInventoryView RemoteVtep
type RemoteVtepInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"vtepIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"poolUuid,omitempty"`
}

