// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdatePciDeviceParamDetail UpdatePciDevice detail param
type UpdatePciDeviceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MetaData string `json:"metaData,omitempty"`
}

// UpdatePciDeviceParam UpdatePciDevice request param
type UpdatePciDeviceParam struct {
	BaseParam
	UpdatePciDevice UpdatePciDeviceParamDetail `json:"updatePciDevice"`
}
// DeletePciDeviceParamDetail DeletePciDevice detail param
type DeletePciDeviceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceParam DeletePciDevice request param
type DeletePciDeviceParam struct {
	BaseParam
	DeletePciDevice DeletePciDeviceParamDetail `json:"deletePciDevice"`
}
