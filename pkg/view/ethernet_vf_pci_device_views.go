// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EthernetVfPciDeviceInventoryView EthernetVfPciDevice
type EthernetVfPciDeviceInventoryView struct {
	rest string `json:"hostDevUuid,omitempty"`
	rest string `json:"interfaceName,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"vfStatus,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"pciSpecUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"virtStatus,omitempty"`
	rest string `json:"chooser,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"device,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"pciDeviceAddress,omitempty"`
	rest string `json:"iommuGroup,omitempty"`
	rest interface{} `json:"metaData,omitempty"`
	rest string `json:"rev,omitempty"`
	rest string `json:"dependentDevices,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PciDevicePciDeviceOfferingRefInventoryView `json:"matchedPciDeviceOfferingRef,omitempty"`
	rest []PciDeviceMdevSpecRefInventoryView `json:"mdevSpecRefs,omitempty"`
}

