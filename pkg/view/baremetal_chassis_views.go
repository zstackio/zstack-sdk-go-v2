// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalChassisInventoryView BaremetalChassis
type BaremetalChassisInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"pxeServerUuid,omitempty"`
	rest string `json:"ipmiAddress,omitempty"`
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiUsername,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []BaremetalHardwareInfoInventoryView `json:"hardwareInfos,omitempty"`
}

