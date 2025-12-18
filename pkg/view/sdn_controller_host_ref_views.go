// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SdnControllerHostRefInventoryView SdnControllerHostRef
type SdnControllerHostRefInventoryView struct {
	rest string `json:"sdnControllerUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest string `json:"vtepIp,omitempty"`
	rest string `json:"nicPciAddresses,omitempty"`
	rest string `json:"nicDrivers,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"bondMode,omitempty"`
	rest string `json:"lacpMode,omitempty"`
}

