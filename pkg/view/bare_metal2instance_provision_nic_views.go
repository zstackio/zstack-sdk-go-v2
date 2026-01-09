// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2InstanceProvisionNicInventoryView BareMetal2InstanceProvisionNic
type BareMetal2InstanceProvisionNicInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	NetworkUuid *string `json:"networkUuid,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

