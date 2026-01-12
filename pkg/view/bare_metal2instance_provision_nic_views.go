// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2InstanceProvisionNicInventoryView BareMetal2InstanceProvisionNic
type BareMetal2InstanceProvisionNicInventoryView struct {
	BaseInfoView
	BaseTimeView
	NetworkUuid *string `json:"networkUuid,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
}

