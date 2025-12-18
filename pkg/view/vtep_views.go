// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VtepInventoryView Vtep
type VtepInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vtepIp,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"physicalInterface,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"poolUuid,omitempty"`
}

