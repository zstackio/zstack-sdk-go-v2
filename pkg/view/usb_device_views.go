// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UsbDeviceInventoryView UsbDevice
type UsbDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	State string `json:"state,omitempty"`
	BusNum string `json:"busNum,omitempty"`
	DevNum string `json:"devNum,omitempty"`
	IdVendor string `json:"idVendor,omitempty"`
	IdProduct string `json:"idProduct,omitempty"`
	IManufacturer string `json:"iManufacturer,omitempty"`
	IProduct string `json:"iProduct,omitempty"`
	ISerial string `json:"iSerial,omitempty"`
	UsbVersion string `json:"usbVersion,omitempty"`
	AttachType string `json:"attachType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

