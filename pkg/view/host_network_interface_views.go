// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkInterfaceInventoryView HostNetworkInterface
type HostNetworkInterfaceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"bondingUuid,omitempty"`
	rest string `json:"interfaceModel,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"interfaceName,omitempty"`
	rest string `json:"interfaceType,omitempty"`
	rest int64 `json:"speed,omitempty"`
	rest bool `json:"slaveActive,omitempty"`
	rest bool `json:"carrierActive,omitempty"`
	rest []string `json:"ipAddresses,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"mac,omitempty"`
	rest string `json:"callBackIp,omitempty"`
	rest string `json:"pciDeviceAddress,omitempty"`
	rest string `json:"driverType,omitempty"`
	rest string `json:"offloadStatus,omitempty"`
	rest string `json:"virtStatus,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

