// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeletePciDeviceOfferingParamDetail DeletePciDeviceOffering detail param
type DeletePciDeviceOfferingParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceOfferingParam DeletePciDeviceOffering request param
type DeletePciDeviceOfferingParam struct {
	BaseParam
	Params DeletePciDeviceOfferingParamDetail `json:"params"`
}
// CreatePciDeviceOfferingParamDetail CreatePciDeviceOffering detail param
type CreatePciDeviceOfferingParamDetail struct {
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
	Params CreatePciDeviceOfferingParamDetail `json:"params"`
}
