// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PciDevicePciDeviceOfferingRefInventoryView PciDevicePciDeviceOfferingRef
type PciDevicePciDeviceOfferingRefInventoryView struct {
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid,omitempty"`
}

