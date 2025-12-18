// Copyright (c) ZStack.io, Inc.

package param

// DeletePciDeviceOfferingDetailParam DeletePciDeviceOffering详细参数
type DeletePciDeviceOfferingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceOfferingParam DeletePciDeviceOffering请求参数
type DeletePciDeviceOfferingParam struct {
	BaseParam
	Params DeletePciDeviceOfferingDetailParam `json:"params"` // 详细参数
}

