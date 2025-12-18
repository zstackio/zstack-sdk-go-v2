// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GpuDeviceSpecInventoryView GpuDeviceSpec
type GpuDeviceSpecInventoryView struct {
	rest int64 `json:"memory,omitempty"`
	rest string `json:"gpuType,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vendorId,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"deviceId,omitempty"`
	rest string `json:"device,omitempty"`
	rest string `json:"subvendorId,omitempty"`
	rest string `json:"subdeviceId,omitempty"`
	rest string `json:"ramSize,omitempty"`
	rest int `json:"maxPartNum,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest bool `json:"isVirtual,omitempty"`
	rest bool `json:"allowResourceConfigWithMultipleDevices,omitempty"`
	rest string `json:"romVersion,omitempty"`
	rest string `json:"romMd5sum,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int `json:"maxAvailableDevicesPerHost,omitempty"`
}

