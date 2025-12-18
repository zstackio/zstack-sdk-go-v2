// Copyright (c) ZStack.io, Inc.

package param

// GenerateSeMdevDevicesDetailParam GenerateSeMdevDevices详细参数
type GenerateSeMdevDevicesDetailParam struct {
	rest string `json:"mttyDeviceUuid" validate:"required"` // 必填
	rest int `json:"virtPartNum" validate:"required"` // 必填
}

// GenerateSeMdevDevicesParam GenerateSeMdevDevices请求参数
type GenerateSeMdevDevicesParam struct {
	BaseParam
	Params GenerateSeMdevDevicesDetailParam `json:"params"` // 详细参数
}

