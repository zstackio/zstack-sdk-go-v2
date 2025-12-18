// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UsbDeviceInventoryView UsbDevice
type UsbDeviceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"busNum,omitempty"`
	rest string `json:"devNum,omitempty"`
	rest string `json:"idVendor,omitempty"`
	rest string `json:"idProduct,omitempty"`
	rest string `json:"iManufacturer,omitempty"`
	rest string `json:"iProduct,omitempty"`
	rest string `json:"iSerial,omitempty"`
	rest string `json:"usbVersion,omitempty"`
	rest string `json:"attachType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

