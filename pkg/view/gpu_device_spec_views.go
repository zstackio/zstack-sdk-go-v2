// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GpuDeviceSpecInventoryView GpuDeviceSpec
type GpuDeviceSpecInventoryView struct {
	Memory int64 `json:"memory,omitempty"`
	GpuType string `json:"gpuType,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	Device string `json:"device,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	RamSize string `json:"ramSize,omitempty"`
	MaxPartNum int `json:"maxPartNum,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	IsVirtual bool `json:"isVirtual,omitempty"`
	AllowResourceConfigWithMultipleDevices bool `json:"allowResourceConfigWithMultipleDevices,omitempty"`
	RomVersion string `json:"romVersion,omitempty"`
	RomMd5sum string `json:"romMd5sum,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	MaxAvailableDevicesPerHost int `json:"maxAvailableDevicesPerHost,omitempty"`
}

