// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalInstanceInventoryView BaremetalInstance
type BaremetalInstanceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"pxeServerUuid,omitempty"`
	rest string `json:"chassisUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest string `json:"templateUuid,omitempty"`
	rest string `json:"platform,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"username,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []BaremetalNicInventoryView `json:"bmNics,omitempty"`
}

