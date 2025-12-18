// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VCenterInventoryView VCenter
type VCenterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"domainName,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"userName,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"version,omitempty"`
	rest bool `json:"https,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

