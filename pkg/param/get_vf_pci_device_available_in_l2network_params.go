// Copyright (c) ZStack.io, Inc.

package param

// GetVfPciDeviceAvailableInL2NetworkDetailParam GetVfPciDeviceAvailableInL2Network detail param
type GetVfPciDeviceAvailableInL2NetworkDetailParam struct {
	L2NetworkUuids []string `json:"l2NetworkUuids" validate:"required"`
}

// GetVfPciDeviceAvailableInL2NetworkParam GetVfPciDeviceAvailableInL2Network request param
type GetVfPciDeviceAvailableInL2NetworkParam struct {
	BaseParam
	Params GetVfPciDeviceAvailableInL2NetworkDetailParam `json:"params"`
}
