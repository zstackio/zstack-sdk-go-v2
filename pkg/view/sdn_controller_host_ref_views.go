// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SdnControllerHostRefInventoryView SdnControllerHostRef
type SdnControllerHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	NicPciAddresses string `json:"nicPciAddresses,omitempty"`
	NicDrivers string `json:"nicDrivers,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	BondMode string `json:"bondMode,omitempty"`
	LacpMode string `json:"lacpMode,omitempty"`
}

