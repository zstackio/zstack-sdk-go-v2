// Copyright (c) ZStack.io, Inc.

package param

// UngenerateSriovPciDevicesDetailParam UngenerateSriovPciDevices详细参数
type UngenerateSriovPciDevicesDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
}

// UngenerateSriovPciDevicesParam UngenerateSriovPciDevices请求参数
type UngenerateSriovPciDevicesParam struct {
	BaseParam
	Params UngenerateSriovPciDevicesDetailParam `json:"params"` // 详细参数
}

