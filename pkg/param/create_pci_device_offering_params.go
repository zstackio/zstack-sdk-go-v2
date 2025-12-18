// Copyright (c) ZStack.io, Inc.

package param

// CreatePciDeviceOfferingDetailParam CreatePciDeviceOffering detail param
type CreatePciDeviceOfferingDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VendorId string `json:"vendorId" validate:"required"`
	DeviceId string `json:"deviceId" validate:"required"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	RamSize string `json:"ramSize,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePciDeviceOfferingParam CreatePciDeviceOffering request param
type CreatePciDeviceOfferingParam struct {
	BaseParam
	Params CreatePciDeviceOfferingDetailParam `json:"params"`
}
