// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2GatewayProvisionNicInventoryView BareMetal2GatewayProvisionNic
type BareMetal2GatewayProvisionNicInventoryView struct {
	BaseInfoView
	BaseTimeView
	NetworkUuid string `json:"networkUuid,omitempty"`
	InterfaceName string `json:"interfaceName,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Metadata string `json:"metadata,omitempty"`
}

