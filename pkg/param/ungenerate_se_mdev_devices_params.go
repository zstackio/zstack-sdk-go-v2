// Copyright (c) ZStack.io, Inc.

package param

// UngenerateSeMdevDevicesDetailParam UngenerateSeMdevDevices详细参数
type UngenerateSeMdevDevicesDetailParam struct {
	rest string `json:"mttyDeviceUuid" validate:"required"` // 必填
}

// UngenerateSeMdevDevicesParam UngenerateSeMdevDevices请求参数
type UngenerateSeMdevDevicesParam struct {
	BaseParam
	Params UngenerateSeMdevDevicesDetailParam `json:"params"` // 详细参数
}

