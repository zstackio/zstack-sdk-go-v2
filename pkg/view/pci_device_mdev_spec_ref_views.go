// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PciDeviceMdevSpecRefInventoryView PciDeviceMdevSpecRef
type PciDeviceMdevSpecRefInventoryView struct {
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	Effective bool `json:"effective,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

