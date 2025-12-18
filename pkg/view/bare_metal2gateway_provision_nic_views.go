// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2GatewayProvisionNicInventoryView BareMetal2GatewayProvisionNic
type BareMetal2GatewayProvisionNicInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"networkUuid,omitempty"`
	rest string `json:"interfaceName,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

