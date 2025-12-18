// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalPxeServerInventoryView BaremetalPxeServer
type BaremetalPxeServerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"hostname,omitempty"`
	rest string `json:"sshUsername,omitempty"`
	rest string `json:"sshPassword,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"storagePath,omitempty"`
	rest string `json:"dhcpInterface,omitempty"`
	rest string `json:"dhcpInterfaceAddress,omitempty"`
	rest string `json:"dhcpRangeBegin,omitempty"`
	rest string `json:"dhcpRangeEnd,omitempty"`
	rest string `json:"dhcpRangeNetmask,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int64 `json:"totalCapacity,omitempty"`
	rest int64 `json:"availableCapacity,omitempty"`
	rest []string `json:"attachedClusterUuids,omitempty"`
}

