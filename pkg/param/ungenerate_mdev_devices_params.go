// Copyright (c) ZStack.io, Inc.

package param

// UngenerateMdevDevicesDetailParam UngenerateMdevDevices详细参数
type UngenerateMdevDevicesDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
}

// UngenerateMdevDevicesParam UngenerateMdevDevices请求参数
type UngenerateMdevDevicesParam struct {
	BaseParam
	Params UngenerateMdevDevicesDetailParam `json:"params"` // 详细参数
}

